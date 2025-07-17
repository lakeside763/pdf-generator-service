package service

import "github.com/lakeside763/pdf-generator-service/client"

type StudentService interface {
	GenerateReport(id string) ([]byte, error)
}

type studentService struct {
	client       client.StudentClient
	pdfGenerator PDFGenerator
}

func NewStudentService(client client.StudentClient, pdfGenerator PDFGenerator) StudentService {
	return &studentService{
		client:       client,
		pdfGenerator: pdfGenerator,
	}
}

// GenerateReport implements StudentService.
func (s *studentService) GenerateReport(id string) ([]byte, error) {
	student, err := s.client.FetchStudentByID(id)
	if err != nil {
		return nil, err
	}
	return s.pdfGenerator.Generate(student)
}
