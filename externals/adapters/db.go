package adapters

import (
	"database/sql"
	"notes/utils/config"
	"strconv"
	"time"

	mysql "github.com/go-sql-driver/mysql"
)

func NewDB(conf config.Database) (*sql.DB, error) {

	config := mysql.Config{
		User:                 conf.User,
		Passwd:               conf.Password,
		Net:                  "tcp",
		Addr:                 conf.Host + ":" + strconv.Itoa(conf.Port),
		DBName:               conf.Database,
		AllowNativePasswords: true,
		ParseTime:            true,
		Timeout:              time.Duration(conf.Timeout) * time.Millisecond,
		ReadTimeout:          time.Duration(conf.ReadTimeout) * time.Millisecond,
		WriteTimeout:         time.Duration(conf.WriteTimeout) * time.Millisecond,
	}

	conn := config.FormatDSN()

	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(conf.OpenConnection)
	db.SetMaxIdleConns(conf.IdleConnection)
	db.SetConnMaxLifetime(time.Duration(conf.ConnectionLifeTime) * time.Millisecond)
	db.SetConnMaxIdleTime(time.Duration(conf.ConnectionIdleTime) * time.Millisecond)

	return db, nil
}
