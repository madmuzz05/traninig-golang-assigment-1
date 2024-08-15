package main

import (
	"os"

	student "github.com/madmuzz05/traninig-golang-assigment-1/repository"
)

func main() {

	name := os.Args[1]

	var person = student.Students{}

	students := []student.Students{
		{Name: "Ucil", Address: "Jalan Hati-Hati", JobTitle: "Kepala Sekolah", Reason: "Ingin merubah sampang menjadi modern"},
		{Name: "Yoci", Address: "Jalan terus jadian kagak", JobTitle: "Pesenam", Reason: "Ingin membuat aplikasi senam"},
		{Name: "Uza", Address: "Jalan Buntu", JobTitle: "Wirausaha", Reason: "Ada deh"},
	}

	person.ShowDataByName(name, students)

}
