package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/pdf-generator-service/client"
	"github.com/lakeside763/pdf-generator-service/config"
	"github.com/lakeside763/pdf-generator-service/handler"
	"github.com/lakeside763/pdf-generator-service/service"
)

func main() {
	cfg := config.New()

	studentClient := client.NewStudentClient(cfg.BaseURL)
	pdfGenerator := service.NewPDFGenerator()
	studentService := service.NewStudentService(studentClient, pdfGenerator)

	h := handler.NewStudentHandler(studentService)

	router := httprouter.New()
	router.GET("/api/v1/students/:id/report", h.GenerateStudentReport)

	log.Println("Go PDF microservice is running on :4500")
	log.Fatal(http.ListenAndServe(":4500", router))
}
