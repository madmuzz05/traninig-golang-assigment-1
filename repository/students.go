package repository

import (
	"fmt"
	"strings"
)

type Students struct {
	Id       int
	Name     string
	Address  string
	JobTitle string
	Reason   string
}

func (selectedPerson Students) ShowDataByName(studentName string, students []Students) {

	for i, student := range students {
		if strings.EqualFold(strings.ToLower(studentName), strings.ToLower(student.Name)) {
			selectedPerson.Id = i
			selectedPerson.Name = student.Name
			selectedPerson.Address = student.Address
			selectedPerson.JobTitle = student.JobTitle
			selectedPerson.Reason = student.Reason
		}
	}
	if selectedPerson.Name != "" {
		fmt.Println("ID : ", selectedPerson.Id)
		fmt.Println("nama : ", selectedPerson.Name)
		fmt.Println("alamat : ", selectedPerson.Address)
		fmt.Println("pekerjaan : ", selectedPerson.JobTitle)
		fmt.Println("alasan : ", selectedPerson.Reason)
	}
}
