package models

type (
	// Book structure
	Book struct {
		ID     string `json:"id" bson:"id"`
		Name   string `json:"name" bson:"name"`
		Author string `json:"author" bson:"author"`
	}
	// GetBookRepoReq structure
	GetBookRepoReq struct {
		ID string `json:"id"`
	}
	// GetBookSrvReq structure
	GetBookSrvReq struct {
		ID string `json:"id"`
	}
	// GetBookReq structure
	GetBookReq struct {
		ID string `json:"id"`
	}
	InsertBookReq struct {
		Name   string `json:"name"`
		Author string `json:"author"`
	}
)
