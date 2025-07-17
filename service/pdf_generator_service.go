package service

import (
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/lakeside763/pdf-generator-service/model"
)

type PDFGenerator interface {
	Generate(student *model.Student) ([]byte, error)
}

type pdfGenerator struct {}

func NewPDFGenerator() PDFGenerator {
	return &pdfGenerator{}
}

func (g *pdfGenerator) Generate(student *model.Student) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.Cell(40, 10, fmt.Sprintf("Student Report for %s", student.Name))
	pdf.Ln(12)

	// Dynamically render fields
	fields := map[string]string{
		"Email": student.Email,
		"Phone": student.Phone,
		"Gender": student.Gender,
		"Class": student.Class,
		"Section": student.Section,
		"Roll": student.Roll,
		"Father's Name": student.FatherName,
		"Mother's Name": student.MotherName,
		"Guardian's Name": student.GuardianName,
		"Admission Date": student.AdmissionDate,
	}

	for label, value := range fields {
		pdf.Cell(50, 10, fmt.Sprintf("%s: %s", label, value))
		pdf.Ln(8)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}