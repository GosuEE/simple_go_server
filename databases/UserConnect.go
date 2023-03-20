package databases

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"

	"log"

	"entity"
)

func GetAllUsers() []entity.User{
	db, err := sql.Open("mysql", "testuser:test@tcp(127.0.0.1:3306)/TESTDB")

	if err != nil{
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM user")

	if err != nil{
		log.Fatal(err)
	}

	ret := make([]entity.User,0)

	for rows.Next(){
		u := entity.User{}
		if err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.Job); err != nil {
			fmt.Println(err)
		}
		ret =append(ret,u)
	}
	rows.Close()

	db.Close()

	return ret;
}