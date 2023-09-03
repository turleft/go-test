package src

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
)

type Test struct {
	Id       int    `db:"id"`
	Num      int    `db:"num"`
	UserName string `db:"user_name"`
}

var Db *sqlx.DB
var wg sync.WaitGroup

func init() {
	database, err := sqlx.Open("mysql", "root:al52ben0@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		database.Close()
		return
	}
	Db = database
}
func init() {
	fmt.Println("second init function")
}
func insert(num int) {
	defer wg.Done() // goroutine结束就登记-1
	r, err := Db.Exec("insert into test(num,user_name)values(?,?)", num, num*num)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert succ:", id)

}

func transaction() {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}
	r, err := conn.Exec("insert into test(num,user_name) values (?,?)", "10", "liming")
	if err != nil {
		fmt.Println("insert failed :", err)
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("insert failed :", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert success :", id)

	r, err = conn.Exec("insert into test(num,user_name) values (?,?)", "11", "zhangsan")
	if err != nil {
		fmt.Println("insert failed :", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("insert failed :", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert success :", id)
	conn.Commit()
}

func MyDB() {
	//for i := 0; i < 10; i++ {
	//	_, err := Db.Exec("insert into test(num,user_name)values(?, ?)", i, "name_"+strconv.Itoa(i))
	//	if err != nil {
	//		fmt.Println("exec failed, ", err)
	//		return
	//	}
	//}
	var test []Test
	err := Db.Select(&test, "select * from test where id<?", 10)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	for _, data := range test {
		fmt.Println("select succ:", data)
	}
	transaction()
}
