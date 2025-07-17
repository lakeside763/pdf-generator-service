package service

import (
	"errors"
	"testing"

	"github.com/lakeside763/pdf-generator-service/model"
	"github.com/stretchr/testify/assert"
)

// --- Mock Client ---
type mockClient struct {
	FetchFunc func(id string) (*model.Student, error)
}

func (m *mockClient) FetchStudentByID(id string) (*model.Student, error) {
	return m.FetchFunc(id)
}

// --- Mock PDF Generator ---
type mockPDFGenerator struct {
	GenerateFunc func(student *model.Student) ([]byte, error)
}

func (m *mockPDFGenerator) Generate(student *model.Student) ([]byte, error) {
	return m.GenerateFunc(student)
}

func TestGenerateReport_Success(t *testing.T) {
	mockStudent := &model.Student{ID: "123", Name: "Test Student"}
	mockPDF := []byte("pdf content")

	service := NewStudentService(
		&mockClient{
			FetchFunc: func(id string) (*model.Student, error) {
				return mockStudent, nil
			},
		},
		&mockPDFGenerator{
			GenerateFunc: func(student *model.Student) ([]byte, error) {
				return mockPDF, nil
			},
		},
	)

	result, err := service.GenerateReport("123")

	assert.NoError(t, err)
	assert.Equal(t, mockPDF, result)
}

func TestGenerateReport_ClientError(t *testing.T) {
	service := NewStudentService(
		&mockClient{
			FetchFunc: func(id string) (*model.Student, error) {
				return nil, errors.New("client error")
			},
		},
		&mockPDFGenerator{},
	)

	result, err := service.GenerateReport("123")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "client error", err.Error())
}

func TestGenerateReport_PDFGeneratorError(t *testing.T) {
	mockStudent := &model.Student{ID: "123", Name: "Test Student"}

	service := NewStudentService(
		&mockClient{
			FetchFunc: func(id string) (*model.Student, error) {
				return mockStudent, nil
			},
		},
		&mockPDFGenerator{
			GenerateFunc: func(student *model.Student) ([]byte, error) {
				return nil, errors.New("pdf generation error")
			},
		},
	)

	result, err := service.GenerateReport("123")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "pdf generation error", err.Error())
}
