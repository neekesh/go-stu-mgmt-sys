package models

type Student struct {
	Id       uint64 `json:"id,string" gorm:"primaryKey"`
	FullName string `json:"full_name"`
	Age      uint8  `json:"age,string"`
	Address  string `json:"address"`
	Major    string `json:"major"`
	Grade    string `json:"grade"`
}
