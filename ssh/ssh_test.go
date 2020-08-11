package ssh

import (
	"fmt"
	"testing"
	"time"
)


func TestSsh(t *testing.T) {
	conn, err := Connect("", "", "", 22)
	defer conn.Close()
	if err != nil {
		fmt.Println("error:", err)
	}
	cmd := []string {
		"cd /bin",
		"ls -al",
		"exit",
	}
	runRet, err := Run(conn, cmd, 20 * time.Millisecond)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(runRet)
}
