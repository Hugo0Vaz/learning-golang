package src

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
	DBType   string
}

type MySqlDBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type DBConnectionError struct {
	ErrorType    string
	ErrorMessage string
}

func (e *DBConnectionError) Error() string {
	return e.ErrorType + e.ErrorMessage
}

func NewDBConnection(config DBConfig) (*sql.DB, error) {

	switch config.DBType {
	case "mysql":
		return newMysqlConnection(MySqlDBConfig{Host: config.Host,
			Port:     config.Port,
			User:     config.User,
			Pass:     config.Pass,
			Database: config.Database,
		})
	default:
		errorMessage := "the database" + config.DBType + "is not implemented in the tool"
		return nil, &DBConnectionError{ErrorType: "NO DB CLIENT CONFIGURED:", ErrorMessage: errorMessage}
	}

}

func newMysqlConnection(config MySqlDBConfig) (*sql.DB, error) {

	cfg := mysql.Config{
		User:   config.User,
		Passwd: config.Pass,
		Net:    "tcp",
		Addr:   config.Host + config.Port,
		DBName: config.Database,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, err
	}

	return db, nil
}
