package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gopheracademy/slideshow/actions"
	"github.com/gopheracademy/slideshow/models"
	"github.com/markbates/going/defaults"
)

func main() {
	port := defaults.String(os.Getenv("PORT"), "3000")
	baseURL := defaults.String(os.Getenv("CMS_URL"), "https://cms.gopheracademy.com")
	models.BaseURL = baseURL
	log.Printf("Starting slideshow on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), actions.App()))
}
