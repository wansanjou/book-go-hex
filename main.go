package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"wansanjou/handler"
	"wansanjou/logs"
	"wansanjou/repository"
	"wansanjou/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	app := fiber.New()
	db := InitializeDB()

	bookRepository := repository.NewBookRepositoryDB(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	authorRepository := repository.NewAuthorRepostioryDB(db)
	authorService := service.NewAuthorService(authorRepository)
	authorHandler := handler.NewAuthorHandler(authorService)

	publisherRepository := repository.NewPublisherRepositoryDB(db)
	publisherService := service.NewPublisherService(publisherRepository)
	publisherHandler := handler.NewPublisherHandler(publisherService)

	userRepository := repository.NewUserRepositoryFB(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	app.Get("/books",bookHandler.GetBookAll)
	app.Get("/books/:id",bookHandler.GetBookByID)
	app.Post("/books",bookHandler.CreateBook)
	app.Put("/books/:id",bookHandler.UpdateBook)
	app.Delete("/books/:id",bookHandler.DeleteBook)

	app.Get("/authors",authorHandler.GetAuthorAll)
	app.Get("/authors/:id",authorHandler.GetAuthorByID)
	app.Post("/authors",authorHandler.CreateAuthor)
	app.Put("/authors/:id",authorHandler.UpdateAuthor)
	app.Delete("/authors/:id",authorHandler.DeleteAuthor)

	app.Get("/publishers",publisherHandler.GetPublisherAll)
	app.Get("/publishers/:id",publisherHandler.GetPublisherByID)
	app.Post("/publishers",publisherHandler.CreatePublisher)
	app.Put("/publishers/:id",publisherHandler.UpdatePublisher)
	app.Delete("/publishers/:id",publisherHandler.DeletePublisher)

	app.Get("/users",userHandler.GetUserAll)
	app.Get("/users/:id",userHandler.GetUserByID)
	app.Post("/users",userHandler.CreateUser)
	app.Put("/users/:id",userHandler.UpdateUser)
	app.Delete("/users/:id",userHandler.DeleteUser)

	app.Listen(":8080")
}

func InitializeDB() *gorm.DB {
	const (
		host     = "localhost"  // or the Docker service name if running in another container
		port     = 5432         // default PostgreSQL port
		user     = "myuser"     // as defined in docker-compose.yml
		password = "mypassword" // as defined in docker-compose.yml
		dbname   = "mydatabase" // as defined in docker-compose.yml
	)
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	newLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
    logger.Config{
      SlowThreshold:              time.Second,   // Slow SQL threshold
      LogLevel:                   logger.Info,  // Log level
      Colorful:                   true,          // Disable color
    },
  )
	
	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, 
	})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&repository.Book{})
	db.AutoMigrate(&repository.Author{})
	db.AutoMigrate(&repository.Publisher{})
	db.AutoMigrate(&repository.User{})

	logs.Info("Starting server at port :5050")
	return db
}