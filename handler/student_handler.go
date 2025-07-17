package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/pdf-generator-service/service"
)

type StudentHandler struct {
	service service.StudentService
}

func NewStudentHandler(service service.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) GenerateStudentReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	if id == "" {
		http.Error(w, "Missing or invalid student id", http.StatusBadRequest)
		return
	}

	pdfBytes, err := h.service.GenerateReport(id)
	if err != nil {
		http.Error(w, "Failed to generate report: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=student_report.pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)
}
