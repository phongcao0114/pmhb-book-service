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
)

//type (
//	// InsertTransactionRepoReq structure
//	InsertTransactionRepoReq struct {
//		TransactionName string `json:"transaction_name"`
//	}
//	// InsertTransactionSrvReq structure
//	InsertTransactionSrvReq struct {
//		TransactionName string `json:"transaction_name"`
//	}
//	// InsertTransactionReq structure
//	InsertTransactionReq struct {
//		TransactionName string `json:"transaction_name"`
//	}
//	// InsertTransactionSrvRes structure
//	InsertTransactionSrvRes struct {
//		TransactionID   int64  `json:"transaction_id"`
//		TransactionName string `json:"transaction_name"`
//	}
//)
