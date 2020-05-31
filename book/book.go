package book

import(
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"../database"
	"time"
)

type OwnModel struct {
	gorm.Model
    ID        uint       `gorm:"primary_key"`
    CreatedAt time.Time  `json:"-"`
    UpdatedAt time.Time  `json:"-"`
    DeletedAt *time.Time `json:"-"`
}
type Book struct {
	ID        uint       `gorm:"primary_key"`
    CreatedAt time.Time  `json:"-"`
    UpdatedAt time.Time  `json:"-"`
    DeletedAt *time.Time `json:"-";sql:"index"`
	Title string `json:"title"`
	Author string `json:"author"`
	Rating int `json:"rating"`
}

func GetBooks(c *fiber.Ctx){
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
	
}

func GetBook(c *fiber.Ctx){
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func PostBook(c *fiber.Ctx){
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil{
		c.Status(503).Send(err)
	}
	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx){
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("Book not found")
		return
	}
	db.Delete(&book)
	c.Send("Book Deleted")
}