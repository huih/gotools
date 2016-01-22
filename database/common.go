package database
import (
	"database/sql"
	_"github.com/lib/pq"
	"github.com/gotools/logs"
	
)

type ConnInfo struct {
	Username string
	Password string
	Host string
	Port string
	Dbname string
	Sslmode string
	Charset string
}

type DBOper interface {
	ConnDB() (*sql.DB, error)
}

func OpenDB(driverName string, connstr string) (*sql.DB, error) {
	return sql.Open(driverName, connstr)
}