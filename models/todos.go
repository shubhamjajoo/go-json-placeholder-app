package models

type Todo struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserId    int64  `json:"userId"`
}
