package pgdb

import (
	"fmt"
	"errors"
	"database/sql"
	"github.com/gotools/logs"
	_"github.com/lib/pq"
)

var (
	DriveName = "postgres"
)

type PGDB struct {
	DBInfo *ConnInfo
	
	db *sql.DB
}

func (db *PGDB) GenConnStr() (string, error) {
	if len(db.DBInfo.Username) <= 0 {
		return "", errors.New("username is empty")
	}
	
	if len(db.DBInfo.Password) <= 0 {
		return "", errors.New("password is empty")
	}
	
	if len(db.DBInfo.Host) <= 0 {
		cInfo.Host = "127.0.0.1"
		logs.Debug("host is empty, default use 127.0.0.1")
	}
	
	if len(db.DBInfo.Port) <= 0 {
		cInfo.Port = "5432"
		logs.Debug("port is empty, default use 5432")
	}
	
	if len(db.DBInfo.Dbname) <= 0 {
		cInfo.Dbname = cInfo.Username
		logs.Debug("dbname is empty, defualt use user name: %s", cInfo.Username)
	}
	
	if len(db.DBInfo.Sslmode) <= 0 {
		cInfo.Sslmode = "disable"
		logs.Debug("sslmode is empty, default use disable")
	}
	
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", db.DBInfo.Username, 
		db.DBInfo.Password, db.DBInfo.Host, db.DBInfo.Port, db.DBInfo.Dbname, db.DBInfo.Sslmode)
	return connStr
}

func (db *PGDB) ConnDB() (*sql.DB, error) {
	//1. check connection parameter 
	connstr, err := db.GenConnStr()
	if err != nil {
		logs.Error("gen conn str error, err:%s", err.Error())
		return nil, err
	}
	
	db.db, err := OpenDB(DriveName, connstr)
	if err != nil {
		logs.Error("open database error, err: %s", err.Error())
		return nil, err
	}
	
	return db.db, err
}

func (db *PGDB) ReadData(sql string) (rows *sql.Rows, error) {
	rows, err := db.db.Query(sql)
	if err != nil {
		return nil, err
	}
	
	return rows, nil
}

func (db *PGDB) WriteData(sql string) (err error) {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	
	defer db.Close()
	_, err = db.Exec(sql)
	
	return err
}