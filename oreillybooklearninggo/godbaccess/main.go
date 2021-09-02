package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB


type Employee struct {
	emp_no int
	birth_date time.Time
	first_name string
	last_name string
	// gender byte
}
func main() {
	fmt.Println("starting access to database from go lang")
    // Capture connection properties.

	cfg := mysql.Config{
		User: os.Getenv("MARIADB_USER"),
		Passwd: os.Getenv("MARIADB_PASS"),
		Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "employees",
		AllowNativePasswords: true,
		ParseTime: true,
	}



// // Get a driver-specific connector.
// connector, err := mysql.NewConnector(&cfg)
// if err != nil {
//     log.Fatal(err)
// }

// // Get a database handle.
// db = sql.OpenDB(connector)
	    // Get a database handle.

		var err error
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Fatal(err)
		}
		pingErr := db.Ping()

		if pingErr != nil {
			log.Fatal(pingErr)
		}
		fmt.Println(cfg)
		fmt.Println("Connect to database")

		employees, err:= getAllEmployees()
		if err != nil {
			fmt.Println(employees)
		} 
		
			fmt.Println(err)
}

func getAllEmployees() ( []Employee, error){
	fmt.Println("Going to get the employees from table")
	var employees []Employee

	rows, err := db.Query("select emp_no, birth_date, first_name, last_name from employees")
	fmt.Println(rows)
	if err != nil {
        return nil, fmt.Errorf("employees not found or error: %v", err)
    }
	defer rows.Close()
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.emp_no, &emp.birth_date, &emp.first_name, &emp.last_name); err != nil {
			return nil, fmt.Errorf("Employee not found or error: %v", err)
		}
		fmt.Println("found employee %v", emp)
		employees = append(employees, emp)
	}

	if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Employee: %v", err)
    }
    return employees, nil
}