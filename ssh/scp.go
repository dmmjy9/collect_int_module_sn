package ssh

import (
	"bufio"
	"collect_int_module_sn/log"
	"fmt"
	"github.com/tmc/scp"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

func OpenScpServ() {
	Logger := log.InitLogger("openscp.log", "INFO")
	var ips []string
	cmd := []string{
		"system-view",
		"scp server enable",
		"save force",
	}

	ipFile, err := os.Open("ip.list")
	if err != nil {
		Logger.Error("Read ip list file fail")
		return
	}
	defer ipFile.Close()
	line := bufio.NewReader(ipFile)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF { break }
		ips = append(ips, string(content))
	}
	for _, ip := range ips {
		func () {
			conn, err := Connect(ip, "", "", 22)
			defer fmt.Println("run success: " + string(ip))
			if err != nil {
				Logger.Warn("login device fail :" + string(ip))
				fmt.Println("login device fail :" + string(ip))
				return
			}
			runRet, err := Run(conn, cmd, 2)
			Logger.Info(string(ip) + string(runRet))
			Close(conn)
		}()
	}
}

func UploadFile(filename string) {
	Logger := log.InitLogger("upload.log", "INFO")
	var ips []string

	ipFile, err := os.Open("ip.list")
	if err != nil {
		Logger.Error("Read ip list file fail")
		return
	}
	defer ipFile.Close()
	line := bufio.NewReader(ipFile)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF { break }
		ips = append(ips, string(content))
	}


	for _, ip := range ips {
		func() {
			f, err := os.Open(filename)
			if err != nil {
				Logger.Error("Open file error, process exit")
				fmt.Println("Open file error, process exit")
				return
			}
			fileInfo, err := f.Stat()
			if err != nil {
				Logger.Error("Wrong file format, process exit")
				fmt.Println("Wrong file format, process exit")
				return
			}
			defer f.Close()

			conn, err := Connect(ip, "", "", 22)
			session, err := conn.NewSession()
			if err != nil {
				Logger.Warn("Open scp session fail: " + string(ip))
				fmt.Println("Open scp session fail: " + string(ip))
				return
			}
			if err != nil {
				Logger.Warn("Login device fail: " + string(ip))
				fmt.Println("Login fail: " + string(ip))
				return
			}
			defer Close(conn)
			defer session.Close()

			err = scp.Copy(fileInfo.Size(), fileInfo.Mode().Perm(), filename, f, filename, session)
			if err != nil {
				Logger.Warn("Upload file fail: " + string(ip))
				fmt.Println("Upload file fail: " + string(ip))
				return
			}
			Logger.Info("Upload success: " + string(ip))
			fmt.Println("Upload success: " + string(ip))
		}()
	}
}

func Patch(filename string, wg *sync.WaitGroup) {
	Logger := log.InitLogger("patch.log", "INFO")
	var ips []string

	ipFile, err := os.Open("ip.list")
	if err != nil {
		Logger.Error("Read ip list file fail")
		return
	}
	defer ipFile.Close()
	line := bufio.NewReader(ipFile)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF { break }
		ips = append(ips, string(content))
	}

	cmd := []string{
		"install activate patch flash:/" + filename + " all",
		"y",
		"install commit",
		"y",
		"save force",
		"quit",
	}

	for _, ip := range ips {
		wg.Add(1)
		go func(ip string, wg *sync.WaitGroup) {
			defer wg.Done()
			conn, err := Connect(ip, "", "", 22)
			defer Close(conn)
			defer fmt.Println("run success: " + string(ip))
			if err != nil {
				Logger.Warn("login device failed " + string(ip))
				fmt.Println("run fail: " + string(ip))
				return
			}
			runRet, err := Run(conn, cmd, 120*time.Second)
			Logger.Info(string(ip) + string(runRet))
		}(ip, wg)
	}
	wg.Wait()
}

func Check() {
	Logger := log.InitLogger("check.log", "INFO")
	var ips []string
	cmd := []string{
		"screen-length disable",
		"dis version",
		"save force",
		"quit",
	}

	ipFile, err := os.Open("ip.list")
	if err != nil {
		Logger.Error("Read ip list file fail")
		return
	}
	defer ipFile.Close()
	line := bufio.NewReader(ipFile)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF { break }
		ips = append(ips, string(content))
	}
	for _, ip := range ips {
		func () {
			conn, err := Connect(ip, "", "", 22)
			defer Close(conn)
			defer fmt.Println("run success: " + string(ip))
			if err != nil {
				Logger.Warn("login device failed " + string(ip))
				fmt.Println("run fail: " + string(ip))
				return
			}
			runRet, err := Run(conn, cmd, 2)
			if strings.Contains(runRet, "S6800-CMW710-SYSTEM-R2612P02H27.bin") {
				Logger.Info(string(ip) + " success")
			} else {
				Logger.Warn(string(ip) + " fail")
			}
		}()
	}
}
