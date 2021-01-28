package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Book struct {
	gorm.Model
	Title string
	Author string `gorm:"default:'nobody'"`
	Publish int `gorm:"default:2018"`
}

func (b *Book) AfterUpdate(tx *gorm.DB) (err error) {
    fmt.Println("after change:", b)
    return
}

func main() {
	r := gin.Default()

	//dbDSN := "user=gorm password=gorm dbname=database port=5432 sslmode=disable"
	dbDSN := "user=jimu password=psql123456 dbname=jimu_pgms_database port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", dbDSN)
	if err != nil {
		fmt.Println("~ ~ ~ connection failed:", err)
	}
	defer db.Close()

	db.AutoMigrate(&Book{})
	bk1 := Book{Title:"TITLE1", Author:"author1", Publish:2020}
	bk2 := Book{Title:"TITLE2", Publish:2021}
	bk3 := Book{}
	db.Create(&bk1)
	db.Create(&bk2)
	db.Create(&Book{})
	fmt.Println(bk3)
	db.Model(&bk1).Update("title", "TITLE111")
	db.Where("publish = ?", 2021).Delete(&Book{})
	db.Delete(bk2)

	r.Run(":9090")
}