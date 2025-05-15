package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func SettingUpHeaderCsvFile() {
	f, err := os.Create("students.csv")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer f.Close()

	// 2) Create a CSV writer
	writer := csv.NewWriter(f)
	defer writer.Flush() // make sure all buffered data is written

	// 3) Write the header row
	header := []string{
		"StudentName",
		"StudentCPI",
		"StudentRollNo",
		"StudentEmail",
		"StudentDepartment",
		"CompanyName",
		"CompanyRoleOffered",
		"CompanyPrepStrategy",
		"CompanySelectionHighlights",
		"CompanyTips",
	}
	err = writer.Write(header)

	if err != nil {
		log.Fatalf("error writing header to csv: %v", err)
	}
}

func MakingListOfStudents(s StudentInfo) {
	var CompleteStudentInfo []string
	CompleteStudentInfo = append(CompleteStudentInfo,
		s.StudentName,
		strconv.FormatFloat(s.StudentCpi, 'f', 2, 64),
		strconv.Itoa(s.StudentRollNo),
		s.StudentEmail,
		s.StudentDepartment,
		s.CompanyName,
		s.CompanyRoleOffered,
		s.CompanyPrepStrategy,
		s.CompanySelectionHighlights,
		s.CompanyTips,
	)

	file, err := os.OpenFile("students.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("couldn't ope the file %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(CompleteStudentInfo)
	if err != nil {
		log.Fatalf("Couldn't append the file %v", err)
	}

	log.Println("new student data added")
}
