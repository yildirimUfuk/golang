package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	//connect to a database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=postgres user=postgres password=abuzer123")
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to connect to db: %s\n", err))
	}

	defer conn.Close()
	log.Println("connected to database")

	//test my connection
	err = conn.Ping()

	if err != nil {
		log.Fatal("cannot ping database! err is: ", err)
	}

	//get rows from table
	getAllRows(conn)
}

func getAllRows(conn *sql.DB) error {
	fmt.Println("getAllRows method")

	rows, err := conn.Query("select * from userTable")
	if err != nil {
		return err
	}
	defer rows.Close()
	var name, password string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &name, &password)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("record is ", id, name, password)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("error scanning rows ", err)
	}
	fmt.Println("-----------------------------------------------")
	return nil
}
