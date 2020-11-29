package model

type User struct {
	UserId       string    `json:"idUser"`
	IdCardNumber string    `json:"idCardNumber"`
	Username     string    `json:"username"`
	DateOfBirth  string    `json:"dateOfBirth"`
	Job          Job       `json:"job"`
	Education    Education `json:"education"`
	UserStatus   int       `json:"userStatus"`
	CreatedDate  string    `json:"createdDate"`
	UpdatedDate  string    `json:"updatedDate"`
}

type UserList struct {
	Users []*User `json:"users"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	CurrentPage int `json:"currentPage"`
	FirstPage   int `json:"firstPage"`
	LastPage    int `json:"lastPage"`
	TotalData   int `json:"totalData"`
}

type Job struct {
	JobId    string `json:"jobId"`
	JobLabel string `json:"jobLabel"`
}

type Education struct {
	EducationId    string `json:"educationId"`
	EducationLabel string `json:"educationLabel"`
}
