package mysqldb
import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/gotools/logs"
)

var (
	DriveName = "mysql"
)

type MysqlDB struct {
	DBInfo *ConnInfo
}

func (db *MysqlDB) GenConnStr() string {
	if len(db.DBInfo.Username) <= 0 {
		return nil, errors.New("username is empty")
	}
	
	if len(db.DBInfo.Host) <= 0 {
		db.DBInfo.Host = "127.0.0.1"
		logs.Debug("host is empty, default use 127.0.0.1")
	}
	
	if len(db.DBInfo.Port) <= 0 {
		db.DBInfo.Port = "3306"
		logs.Debug("port is empty, default use 3306")
	}
	
	if len(db.DBInfo.Dbname) <= 0 {
		db.DBInfo.Dbname = cInfo.Username
		logs.Debug("dbname is empty, defualt use user name: %s", cInfo.Username)
	}
	
	if (len(db.DBInfo.charset) <= 0) {
		db.DBInfo.Charset = "utf8"
		logs.Debug("charset is empty, default use utf8")
	}
		
	connStr := fmt.Sprintf("%s:%s@%s:%s/%s?charset=%s", db.DBInfo.Username, db.DBInfo.Password, 
		db.DBInfo.Host, db.DBInfo.Port, db.DBInfo.Dbname, db.DBInfo.Charset)
	
	return connStr
}

func (db *MysqlDB) ConnDB() (*sql.DB, error) {
	//1. check connection parameter 
	connstr := db.GenConnStr()
	return OpenDB(DriveName, connStr)
}

