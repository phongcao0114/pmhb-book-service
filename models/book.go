package models

type (
	Book struct {
		ID     string `json:"id" bson:"id"`
		Name   string `json:"name" bson:"name"`
		Author string `json:"author" bson:"author"`
	}
	InsertBookReq struct {
		Name   string `json:"name"`
		Author string `json:"author"`
	}
	UpdateBookReq struct {
		Name   string `json:"name"`
		Author string `json:"author"`
	}
	DeleteBookReq struct {
		ID string `json:"id"`
	}
)
