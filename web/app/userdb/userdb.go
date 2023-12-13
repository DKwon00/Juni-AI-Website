package userdb

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Handler(c *gin.Context) {
	//initialize db
	db, err := sqlx.Open("postgres", "user=postgres password=3348 dbname=Juni-AI-Users sslmode=disable")
	if err != nil {
		log.Fatal(err, "tests")
	}

	rows, err := db.Query("SELECT * FROM juniai_userdb")
	if err != nil {
		log.Fatal(err, "query")
	}

	for rows.Next() {
		var id int
		var username string
		var email string
		err = rows.Scan(&id, &username, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, username, email)
	}
	defer rows.Close()
	// ...
}
