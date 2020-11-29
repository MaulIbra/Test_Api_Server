package model

import (
	"regexp"
	"strings"
)

type User struct {
	UserId       string            `json:"idUser,omitempty"`
	IdCardNumber string            `json:"idCardNumber,omitempty"`
	Username     string            `json:"username,omitempty"`
	DateOfBirth  string            `json:"dateOfBirth,omitempty"`
	Job          Job               `json:"job,omitempty"`
	Education    Education         `json:"education,omitempty"`
	UserStatus   int               `json:"userStatus,omitempty"`
	CreatedDate  string            `json:"createdDate,omitempty"`
	UpdatedDate  string            `json:"updatedDate,omitempty"`
	Errors       map[string]string `json:"errors,omitempty"`
}

func (user *User) Validate() bool {
	user.Errors = make(map[string]string)
	re := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])")
	if len(strings.TrimSpace(user.IdCardNumber)) > 16 || strings.TrimSpace(user.IdCardNumber) == "" {
		user.Errors["IdCardNumber"] = "Please enter a valid id card number"
	}
	if strings.TrimSpace(user.Username) == "" {
		user.Errors["Username"] = "Please enter a username"
	}
	if strings.TrimSpace(user.DateOfBirth) == "" || !re.MatchString(user.DateOfBirth) {
		user.Errors["DateOfBirth"] = "Please enter a valid date of birth"
	}
	if strings.TrimSpace(user.Job.JobId) == "" {
		user.Errors["JobId"] = "Please enter a job id"
	}
	if strings.TrimSpace(user.Education.EducationId) == "" {
		user.Errors["EducationId"] = "Please enter a education id"
	}
	return len(user.Errors) == 0
}

type UserList struct {
	Users    []*User  `json:"users,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}

type Metadata struct {
	CurrentPage int `json:"currentPage,omitempty"`
	FirstPage   int `json:"firstPage,omitempty"`
	LastPage    int `json:"lastPage,omitempty"`
	TotalData   int `json:"totalData,omitempty"`
}

type Job struct {
	JobId    string `json:"jobId,omitempty"`
	JobLabel string `json:"jobLabel,omitempty"`
}

type Education struct {
	EducationId    string `json:"educationId,omitempty"`
	EducationLabel string `json:"educationLabel,omitempty"`
}
