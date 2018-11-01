package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"

		"time"
	"github.com/didi/gendry/manager"
		"fmt"
	"github.com/jinzhu/gorm"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
)

var db *sql.DB
var err error

type Book struct{
	Id int`ddb:"id"`
	Uname string `ddb:"uname"`
	Email string `ddb:"email"`
	Content string`ddb:"content"`
	InsertTime time.Time`ddb:"insert_time"`
}

type Book2 struct {
	Id       int64 `gorm:"primary_key"`
	Uname string
	Email string
	Content string
	InsertTime int
}

var (
	host       = "127.0.0.1"
	port       = 3306
	user       = "root"
	password   = "123456"
	dbName     = "test"
	charsetStr = "utf8"
	connetStr  = "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local"
	driverName = "mysql"

	isInit bool = false
)

var odb *gorm.DB


func main()  {

	testGendry()
	//testOrm()
}

func testOrm(){
	realCon := fmt.Sprintf(connetStr, user, password, host, port, dbName, charsetStr)

	var err error

	//批量查询
	var books []Book2;
	odb, err = gorm.Open("mysql", realCon)
	if err != nil {
		panic(err)
	}
	odb.Table("book").Where("id > ?", 1).Offset(1).Limit(10).Find(&books)
	fmt.Println(len(books))
	fmt.Println(books)




}

func testGendry(){
	//使用gendry
	dbName, user, password, host := "test","root","123456","localhost"
	db, err = manager.New(dbName, user, password, host).Set(
		manager.SetCharset("utf8"),
		//manager.SetAllowCleartextPasswords(true),
		manager.SetInterpolateParams(true),
		manager.SetTimeout(1 * time.Second),
		manager.SetReadTimeout(1 * time.Second)).Port(3306).Open(true)
	fmt.Println(db)

	where := map[string]interface{}{
		"id >": 1,
	}
	table := "book"
	selectFields := []string{"uname", "email", }
	cond, values, _ := builder.BuildSelect(table, where, selectFields)
	rows,err := db.Query(cond, values...)
	defer rows.Close()
	fmt.Println(err)
	var books []Book
	scanner.Scan(rows, &books)
	fmt.Println(len(books))
	for _,book1 := range books{
		fmt.Println(book1)
	}


	/*
	//实现查
	rows,err := db.Query("select id,uname,email,content,insert_time from book where id > 1 limit 10")
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	var books []Book
	scanner.Scan(rows, &books)
	for _,book := range books {
		fmt.Println(book)
	}
	fmt.Println(len(books))
	*/


	//实现增加
}