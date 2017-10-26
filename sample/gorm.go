package sample

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/mochizukikotaro/db-operation-sample/model"
)

func SampleGorm() {
	users := model.CreateUsers()

	db, _ := gorm.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	defer db.Close()
	db.Find(&users) // SELECT * FROM users;
	fmt.Println("gorm")
	fmt.Println(users)
}
