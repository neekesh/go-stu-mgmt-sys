package types

type StudentCreate struct {
	FullName string `json:"full_name" binding:"required"`
	Age      uint8  `json:"age,string" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Major    string `json:"major" binding:"required"`
	Grade    string `json:"grade"  binding:"required"`
}

type UpdateStudentInput struct {
	FullName string `json:"full_name"`
	Age      uint8  `json:"age,string"`
	Address  string `json:"address"`
	Major    string `json:"major"`
	Grade    string `json:"grade"`
}
