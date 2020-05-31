package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/sqlite"
	"./book"
	"./database"
)

func helloWorld(c *fiber.Ctx){
	c.Send("Hello Dunia!")
}

func serverRoot(c *fiber.Ctx){
	c.Send("v1-running")
}

func initDB(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3","books.db")
	if err !=nil{
		panic("Failed to connect")
	}

	fmt.Println("DB Connected")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("DB Migrated")

}

func routeBook(app *fiber.App){
	app.Get("/api/v1/book/", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book/", book.PostBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main(){
	app := fiber.New()
	initDB()
	defer database.DBConn.Close()
	routeBook(app)
	app.Get("/",serverRoot)
	app.Listen(3000)
}
