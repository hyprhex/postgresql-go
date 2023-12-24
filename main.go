package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// type User struct {
// 	ID        int
// 	Age       int
// 	FirstName string
// 	LastName  string
// 	Email     string
// }

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

	// Insert record

	// sqlStatement := `
	// INSERT INTO users (age, email, first_name, last_name)
	// VALUES ($1, $2, $3, $4)
	// RETURNING id`
	// id := 0
	// err = db.QueryRow(sqlStatement, 29, "father@gmai.com", "Father", "Seyf").Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println("New record ID is:", id)
	//

	// Update record
	// sqlStatement := `
	// UPDATE users SET first_name = $2, last_name = $3
	// WHERE id = $1`
	// res, err := db.Exec(sqlStatement, 3, "NewFirst", "NewLast")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// count, err := res.RowsAffected()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println(count)

	// Delete
	// sqlStatement := `
	// DELETE FROM users
	// WHERE id = $1;`
	// res, err := db.Exec(sqlStatement, 8)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// count, err := res.RowsAffected()
	// if err != nil  {
	// 	panic(err)
	// }
	//
	// fmt.Println(count)

	// Querying for a single record
	// sqlStatement := `SELECT id, email FROM users WHERE id = $1;`
	// var email string
	// var id int
	// row := db.QueryRow(sqlStatement, 3)
	// switch err := row.Scan(&id, &email); err {
	// case sql.ErrNoRows:
	// 	fmt.Println("No rows were returend") // Here in real project we return 404 or redirect
	// case nil:
	// 	fmt.Println(id, email)
	// default:
	// 	panic(err) // In real project we can return 500
	// }

	// Return entire record
	// sqlStatement := `SELECT * FROM users WHERE id = $1;`
	// var users User
	// row := db.QueryRow(sqlStatement, 1)
	// errs := row.Scan(&users.ID, &users.Age, &users.FirstName, &users.LastName, &users.Email)
	// switch errs {
	// case sql.ErrNoRows:
	// 	fmt.Println("No rows were returned")
	// 	return
	// case nil:
	// 	fmt.Println(users)
	// default:
	// 	panic(errs)
	// }

	// Querying many records
	rows, err := db.Query("SELECT id, first_name FROM users LIMIT $1", 3)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var firstName string
		err = rows.Scan(&id, &firstName)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, firstName)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

}
