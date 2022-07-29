// connect to a AWS RDS Auroa database, run a query and test network latency to and from the database.
// take in a endpoint to test latency to and from.
// Read the ~/.my.cnf file to get the database credentials.
// Language: go
// Path: main.go
// latency function
// Usage:
// go run main.go -e endpoint -f filename.sql
//

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// global variables
var (
	db  *sql.DB
	err error
)

// flag variables
var (
	endPoint = flag.String("e", "", "endpoint to test latency to and from")
	sqlFile  = flag.String("f", "", "sql file to run")
)

// read the ~/.my.cnf file to get the database credentials
func readMyCnf() {
	file, err := ioutil.ReadFile(os.Getenv("HOME") + "/.my.cnf")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "user") {
			os.Setenv("MYSQL_USER", strings.TrimSpace(line[5:]))
		}
		if strings.HasPrefix(line, "password") {
			os.Setenv("MYSQL_PASSWORD", strings.TrimSpace(line[9:]))
		}
	}
}

// define the latency function
func testLatency() {
	start := time.Now()
	runQuery()
	end := time.Now()
	latency := end.Sub(start)
	fmt.Println("Latency: ", latency)
}

// connect to the database using the go-sql-driver/mysql package
// run the query on the database
func connectToDatabase() {
	db, err = sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+*endPoint+")/")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// reaq from sqlFile and run the query provided in the sqlFile file on the database and print the results of the query ran
func runQuery() {
	file, err := ioutil.ReadFile(*sqlFile)
	if err != nil {
		log.Fatal(err)
	}
	query := string(file)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("--------------------------------------------------------------------------------------------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

// main function to run the program
func main() {
	// call the flag package to parse the command line arguments
	flag.Parse()
	// make sure the endpoint and sqlFile are set
	if *endPoint == "" || *sqlFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	// read the ~/.my.cnf file to get the database credentials. check that the file exists
	if _, err := os.Stat(os.Getenv("HOME") + "/.my.cnf"); os.IsNotExist(err) {
		fmt.Println("Please create a ~/.my.cnf file with the database credentials.")
		os.Exit(1)
	}
	// call the readMyCnf function to read the ~/.my.cnf file
	readMyCnf()
	// connect to the database
	connectToDatabase()
	// run the query on the database and print the results.
	// runQuery()
	// call the latency function to test the latency to and from the database
	testLatency()
	// close the database connection
	db.Close()

}
