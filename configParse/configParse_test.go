package configParse

import (
	"fmt"
	"testing"
)

func TestConfigParse(t *testing.T) {
	if sshconfig, err := GetSSHConfig("config.json"); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(sshconfig.Username)
		fmt.Println(sshconfig.Password)
	}

	if dbconfig, err := GetDBConfig("config.json"); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(dbconfig.Username)
		fmt.Println(dbconfig.Password)
		fmt.Println(dbconfig.Port)
		fmt.Println(dbconfig.Database)
	}
}
