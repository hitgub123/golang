package dbpackage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "yourpassword"
	dbname   = "study"
)

// func main() {
// 	db:=GetDB()
// 	defer db.Close()
// 	Select(db)
// 	// Insert(db)
// 	// Delete(db)
// 	// Update(db)
// }

func CheckError(err error){
	if err != nil{
		panic(err)
	}
}
func Select(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users where id in (3,44,45,46,47,48,49)")
	CheckError(err)
	var es []Product
	for rows.Next() {
		var e Product
		rows.Scan(&e.ID, &e.nAMe, &e.Age)
		es = append(es, e)
	}
	fmt.Printf("%v", es)
}

func Insert(db *sql.DB) {
	stmt,err := db.Prepare(`insert into users(name,age) values($1, $2)`)
	CheckError(err)
    defer stmt.Close()	
	_,err =stmt.Exec("操作系统","666a啊")
	CheckError(err)
}

func Delete(db *sql.DB) {
	stmt,err := db.Prepare(`delete from users where id=$1`)
	CheckError(err)
    defer stmt.Close()	
	_,err =stmt.Exec(43)
	CheckError(err)
}

func Update(db *sql.DB) {
	stmt,err := db.Prepare(`update users set name=$1,age=$2 where id=$3`)
	CheckError(err)
    defer stmt.Close()	
	_,err =stmt.Exec("apple",99,45)
	CheckError(err)
}

type Product struct {
	ID   int8
	nAMe string
	Age  int8
}

func GetDB() (db *sql.DB){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// fmt.Println("Successfully connected!")
	return
}
