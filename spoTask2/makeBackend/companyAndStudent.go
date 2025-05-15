package main

type StudentInfo struct {
	StudentName                string  `json:"studentName"`
	StudentRollNo              int     `json:"studentRollNo"`
	StudentEmail               string  `json:"studentEmail"`
	StudentDepartment          string  `json:"studentDepartment"`
	StudentCpi                 float64 `json:"studentCpi"`
	CompanyName                string  `json:"companyName"`
	CompanyRoleOffered         string  `json:"companyRoleOffered"`
	CompanySelectionHighlights string  `json:"companySelectionHighlights"`
	CompanyPrepStrategy        string  `json:"companyPrepStrategy"`
	CompanyTips                string  `json:"companyTips"`
}

//all the json are made lower case according to best practises know by me
