package database

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var G_db *gorm.DB

type Student struct {
	gorm.Model
	StudentName string
	StudentId   int
	Day         string
	Semester    string
}

type Class struct {
	gorm.Model
	StudentId int
	Location  string
	Day       string
	Lesson    string
	RawWeek   string
	Teacher   string
}

func ConnetDb() {
	db, err := gorm.Open("mysql", "root:mima@/students?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		errors.New("open database failed!")
	}
	G_db = db
}

func CreateTable() {
	if G_db.HasTable(&Student{}) {
		G_db.AutoMigrate()
	} else {
		G_db.CreateTable(&Student{})
	}
	if G_db.HasTable(&Class{}) {
		G_db.AutoMigrate()
	} else {
		G_db.CreateTable(&Class{})
	}
}
