package main

import (
	"fmt"
	"github.com/google/uuid"
)

type Student struct {
	FirstName     string
	LastName      string
	applicationID uuid.UUID
}

func (s Student) ApplicationID() string {
	return s.applicationID.String()
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

func main() {
	uniqueToken, _ := uuid.NewUUID()

	s := Student{
		FirstName:     "Wojtek",
		LastName:      "Olesiński",
		applicationID: uniqueToken,
	}

	// status, err := appdispatcher.Submit(s);
	// logger := log.Default()
	// logger.Println(status, err)
	fmt.Printf("%v", s)
}
