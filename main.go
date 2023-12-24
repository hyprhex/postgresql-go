package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		host   = os.Getenv("HOST")
		port   = os.Getenv("PORT")
		user   = os.Getenv("USERNAME")
		pass   = os.Getenv("PASSWORD")
		dbname = os.Getenv("DBNAME")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")

	sqlStatement := `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, 29, "father@gmai.com", "Father", "Seyf").Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("New record ID is:", id)

}
