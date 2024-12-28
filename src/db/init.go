package init

import (
    "io/ioutil"
    "log"
    "./db.go"
)

func initDB() {
    conn := db.getConnection()
    ddlQuery, err := ioutil.ReadFile("sql/init.sql")
    if err != nil {
        log.Fatalf("Failed to read schema file: %v", err)
    }
    // Execute the SQL
    _, err = conn.Exec(string(ddlQuery))
    if err != nil {
        log.Fatalf("Failed to execute schema: %v", err)
    }
    log.Println("Schema executed successfully!")
}
