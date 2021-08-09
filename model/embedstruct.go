package model

type Base struct {
	B   int
	Tag string
}

type Container struct {
	Base
	Desc string
}

type Job struct {
	// Tạo bảng
	TableName string `json:"table_name" sql:"job.job"`
	// Id
	Id int `json:"id" sql:",pk"`
	// Tên công việc
	Name string `json:"name"`
}
