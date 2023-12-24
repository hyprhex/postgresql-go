package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
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

	fmt.Println(host, port, user, pass, dbname)
}
