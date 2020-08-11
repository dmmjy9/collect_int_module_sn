package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type MySQLDB struct {
	Ipaddr		string
	Port		int
	Username	string
	Password	string
	Database	string
	Dbw			*sql.DB
}

type QueryResult struct {
	Ipaddr		string	`json:"ipaddr"`
	Hostname	string	`json:"hostname"`
	Vendor		string	`json:"vendor"`
	Module		string  `json:"pid"`
}

func (d *MySQLDB) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		d.Username, d.Password, d.Ipaddr, d.Port, d.Database)
	dbw, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	} else {
		d.Dbw = dbw
		return nil
	}
}

func (d *MySQLDB) Query(ipaddr string) (QueryResult, error) {
	qr := QueryResult{}
	rawSql := fmt.Sprintf(
		"SELECT mg_ip, hostname, vendor, pid FROM net_dev_bs_info_copy WHERE mg_ip='%s'", ipaddr)
	err := d.Dbw.QueryRow(rawSql).Scan(&qr.Ipaddr, &qr.Hostname, &qr.Vendor, &qr.Module)
	if err != nil {
		return QueryResult{}, err
	} else {
		qr.Module = strings.Split(qr.Module, "/")[0]
		return qr, nil
	}
}

func (d *MySQLDB) Close() error {
	if err := d.Dbw.Close(); err != nil {
		return err
	} else {
		return nil
	}
}
