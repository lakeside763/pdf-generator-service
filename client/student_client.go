package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/lakeside763/pdf-generator-service/model"
	"github.com/sony/gobreaker"
)

type StudentClient interface {
	FetchStudentByID(id string) (*model.Student, error)
}

type studentClient struct {
	baseURL string
	breaker *gobreaker.CircuitBreaker
}

func NewStudentClient(baseURL string) StudentClient {
	settings := gobreaker.Settings{
		Name:        "StudentClient",
		Timeout:     5 * time.Second,
		MaxRequests: 1,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 3
		},
	}
	return &studentClient{
		baseURL: baseURL,
		breaker: gobreaker.NewCircuitBreaker(settings),
	}
}

// FetchStudentByID fetches a student by ID from the remote service.
func (c *studentClient) FetchStudentByID(id string) (*model.Student, error) {
	result, err := c.breaker.Execute(func() (interface{}, error) {
		url := fmt.Sprintf("%s/students/%s", c.baseURL, id)
		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch student: status %d", resp.StatusCode)
		}

		var student model.Student
		if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
			return nil, err
		}
		return &student, nil
	})

	if err != nil {
		return sampleStudent(), nil // fallback when breaker is open or request fails
	}

	return result.(*model.Student), nil
}

func sampleStudent() *model.Student {
	return &model.Student{
		ID:                 "12345",
		Name:               "John Doe",
		Email:              "john.doe@example.com",
		SystemAccess:       true,
		Phone:              "+1234567890",
		Gender:             "Male",
		DOB:                "2006-05-12",
		Class:              "10",
		Section:            "B",
		Roll:               "32",
		FatherName:         "Robert Doe",
		FatherPhone:        "+1234567891",
		MotherName:         "Jane Doe",
		MotherPhone:        "+1234567892",
		GuardianName:       "Uncle Sam",
		GuardianPhone:      "+1234567893",
		RelationOfGuardian: "Uncle",
		CurrentAddress:     "1234 Elm Street, Springfield",
		PermanentAddress:   "4321 Oak Avenue, Shelbyville",
		AdmissionDate:      "2020-09-01",
		ReporterName:       "Mr. Smith",
	}
}
