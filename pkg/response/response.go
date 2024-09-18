package response

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

var (
	SuccessResponse = Response{Message: "success executing the query", Status: true}
	ErrorResponse   = Response{Message: "failed executing the query", Status: false}
)
