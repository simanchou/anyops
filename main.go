package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Project struct {
	gorm.Model
	Name string `gorm:"not null;unique_index"`
	Type string `gorm:"not null"`
	Apps []App
}

type App struct {
	gorm.Model
	ProjectID int    `gorm:"index"`
	Name      string `gorm:"index"`
	ENV       string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/anyops?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("connect db fail, error: %s", err)
	}
	defer db.Close()

	// auto migrate
	db.AutoMigrate(&Project{}, &App{})

	p0 := Project{
		Name: "sctest01",
		Type: "cp",
		Apps: []App{{Name: "admin", ENV: "beta"}, {Name: "api", ENV: "beta"}, {Name: "pay", ENV: "beta"}},
	}
	db.Create(&p0)

	var p1 Project
	err = db.Preload("Apps").Find(&p1, "1").Error
	if err != nil {
		log.Println(err)
	}
	log.Printf("%v", p1)
}
