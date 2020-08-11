package ssh

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strconv"
	"time"
)

type DataReader struct {
	buffer bytes.Buffer
}

func (s *DataReader) Write(p []byte) (n int, err error) {
	if n, err := s.buffer.Write(p); err != nil {
		return n, err
	} else {
		return len(p), nil
	}
}

func (s *DataReader) Read() string {
	return s.buffer.String()
}

func Connect(ipaddr, username, password string, port int) (*ssh.Client, error) {
	config := &ssh.ClientConfig {
		Config: ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr",
							  "aes128-gcm@openssh.com","arcfour256", "arcfour128",
							  "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},},
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", ipaddr + ":" + strconv.Itoa(port), config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Run(client *ssh.Client, commands []string, waittime time.Duration) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:     0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err = session.RequestPty("xterm",80,100, modes)
	if err != nil {
		return "", err
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		return "", err
	}

	dataOut := new(DataReader)
	dataErr := new(DataReader)
	session.Stdout = dataOut
	session.Stderr = dataErr

	err = session.Shell()
	if err != nil {
		return "", err
	}

	for _, cmd := range commands {
		_, err = fmt.Fprintf(stdin, "%s\n", cmd)
		if err != nil {
			return "", err
		}
		time.Sleep(waittime)
	}

	go func(session *ssh.Session) {
		time.Sleep(10 * time.Second)
		err := session.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(session)

	err = session.Wait()

	dataOutStr := dataOut.Read()
	session.Close()
	return dataOutStr, nil
}

func Close(client *ssh.Client) error {
	err := client.Close()
	if err != nil {
		return err
	}
	return nil
}
