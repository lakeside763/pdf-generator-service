package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// --- Mock StudentService ---
type mockStudentService struct {
	GenerateReportFunc func(id string) ([]byte, error)
}

func (m *mockStudentService) GenerateReport(id string) ([]byte, error) {
	return m.GenerateReportFunc(id)
}

func TestGenerateStudentReport_Success(t *testing.T) {
	mockPDF := []byte("pdf content")

	handler := NewStudentHandler(&mockStudentService{
		GenerateReportFunc: func(id string) ([]byte, error) {
			assert.Equal(t, "12345", id)
			return mockPDF, nil
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/students/12345/report", nil)
	rr := httptest.NewRecorder()

	// simulate httprouter.Params
	params := httprouter.Params{httprouter.Param{Key: "id", Value: "12345"}}

	handler.GenerateStudentReport(rr, req, params)

	resp := rr.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/pdf", resp.Header.Get("Content-Type"))
	assert.Equal(t, "attachment; filename=student_report.pdf", resp.Header.Get("Content-Disposition"))
}
