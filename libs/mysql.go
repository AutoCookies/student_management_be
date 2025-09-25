package libs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// ConnectMySQL kết nối tới MySQL
func ConnectMySQL(user, password, host, dbname string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, host, dbname)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening DB: %w", err)
	}

	// Test kết nối
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("error ping DB: %w", err)
	}

	log.Println("Connected to MySQL")
	return nil
}
