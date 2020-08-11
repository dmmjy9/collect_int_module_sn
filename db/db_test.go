package db

import (
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	db := MySQLDB {
		Ipaddr:		"",
		Port:		8888,
		Username:	"",
		Password:	"",
		Database:	"",
		Dbw:		nil,
	}
	err := db.Connect()
	if err != nil {
		fmt.Println("connect to mysql fail")
	}
	rawRet, err := db.Query("172.29.99.179")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rawRet.Ipaddr, rawRet.Hostname, rawRet.Vendor, rawRet.Module)
	}
	if err := db.Close(); err != nil {
		fmt.Println("close mysql fail")
	}
}
