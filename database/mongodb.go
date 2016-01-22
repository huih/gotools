package mongodb

import (
	"fmt"
	"errors"
	"database/sql"
	"github.com/gotools/logs"
	_"github.com/go-mgo/mgo"
)

var (
	DriveName = "mongodb"
)

type MogoDB struct {
	DBInfo *ConnInfo
}

func (db *MogoDB) GenConnStr() string {
	if len(db.DBInfo.Username) <= 0 {
		return nil, errors.New("username is empty")
	}
	
	if len(db.DBInfo.Host) <= 0 {
		cInfo.Host = "127.0.0.1"
		logs.Debug("host is empty, default use 127.0.0.1")
	}
	
	if len(db.DBInfo.Port) <= 0 {
		cInfo.Port = "27017"
		logs.Debug("port is empty, default use 5432")
	}
	
	if len(db.DBInfo.Dbname) <= 0 {
		cInfo.Dbname = cInfo.Username
		logs.Debug("dbname is empty, defualt use user name: %s", cInfo.Username)
	}
		
	connStr := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", db.DBInfo.Username, 
		db.DBInfo.Password, db.DBInfo.Host, db.DBInfo.Port, db.DBInfo.Dbname)
	return connStr
}

func (db *MogoDB) ConnDB() (*sql.DB, error) {
	//1. check connection parameter 
	connstr := db.GenConnStr()
	return OpenDB(DriveName, connStr)
}