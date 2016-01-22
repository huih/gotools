package oracledb

import (
	"fmt"
	"errors"
	"database/sql"
	"github.com/gotools/logs"
	_"github.com/mattn/go-oci8"
)

var (
	DriveName = "oci8"
)

type OracleDB struct {
	DBInfo *ConnInfo
}

func (db *OracleDB) GenConnStr() (string, error) {
	if len(db.DBInfo.Username) <= 0 {
		return "", errors.New("username is empty")
	}
	
	if len(db.DBInfo.Host) <= 0 {
		cInfo.Host = "127.0.0.1"
		logs.Debug("host is empty, default use 127.0.0.1")
	}
	
	if len(db.DBInfo.Dbname) <= 0 {
		cInfo.Dbname = cInfo.Username
		logs.Error("dbname is empty, error")
		return "", errors.New("dbname is empty")
	}
	
	connStr := fmt.Sprintf("%s/%s@%s", db.DBInfo.Username, db.DBInfo.Password, db.DBInfo.Dbname)
	return connStr, nil
}

func (db *OracleDB) ConnDB() (*sql.DB, error) {
	//1. check connection parameter 
	connstr, err := db.GenConnStr()
	if err != nil {
		return nil, err
	}
	
	return OpenDB(DriveName, connStr)
}