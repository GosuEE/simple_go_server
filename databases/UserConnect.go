package databases

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"

	"log"

	"entity"

	requestdto "requestDTO"
)

func intNull(num int) sql.NullInt32{
	if num == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: int32(num),
		Valid: true,
	}
}

func strNull(str string) sql.NullString{
	if str == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: str,
		Valid: true,
	}
}

func DeleteUser(u *requestdto.UserRequestDTO){
	db, err := sql.Open("mysql", "testuser:test@tcp(127.0.0.1:3306)/TESTDB")

	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM user WHERE id = ?", u.Id)

	if err != nil{
		log.Fatal(err)
	}

	n, err:= result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row deleted");
	}

}

func UpdateUser(u *requestdto.UserRequestDTO){
	db, err := sql.Open("mysql", "testuser:test@tcp(127.0.0.1:3306)/TESTDB")

	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	query :=
`
UPDATE user
	SET name = IFNULL(?, name),
	age = IFNULL(?, age),
	job = IFNULL(?, job)
	where id = ?
`

	result, err := db.Exec(query, strNull(u.Name), intNull(u.Age), strNull(u.Job), u.Id)

	if err != nil{
		log.Fatal(err)
	}

	n, err:= result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row updated");
	}
}

func InsertUser(u *requestdto.UserRequestDTO){
	db, err := sql.Open("mysql", "testuser:test@tcp(127.0.0.1:3306)/TESTDB")

	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO user (name, age, job) VALUES(?, ?, ?)", u.Name, u.Age, u.Job)

	if err != nil{
		log.Fatal(err)
	}

	n, err:= result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted");
	}
}

func GetAllUsers() []entity.UserEntity{
	db, err := sql.Open("mysql", "testuser:test@tcp(127.0.0.1:3306)/TESTDB")

	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user")

	if err != nil{
		log.Fatal(err)
	}

	ret := make([]entity.UserEntity,0)

	for rows.Next(){
		u := entity.UserEntity{}
		if err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.Job); err != nil {
			log.Fatal(err)
		}
		ret =append(ret,u)
	}
	rows.Close()

	db.Close()

	return ret;
}