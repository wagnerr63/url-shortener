package main

import (
	"log"
	"os"
	"url-shortener/configs"
	"url-shortener/controllers"
	"url-shortener/repositories"
	"url-shortener/usecases"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repositories := repositories.New(repositories.Options{
		WriterGorm: configs.GetWriterGorm(),
		ReaderGorm: configs.GetReaderGorm(),
	})

	router := configs.InitTools()

	usecases := usecases.New(usecases.Options{
		Repo: repositories,
	})

	controllers := controllers.New(controllers.Options{Usecases: usecases})

	router.POST("/add", controllers.ShortenedUrl.Create)
	router.GET("/{id}", controllers.ShortenedUrl.FindByShortenedUrl)
	router.GET("/", controllers.ShortenedUrl.Index)
	router.SERVE(":" + os.Getenv("PORT"))
}
