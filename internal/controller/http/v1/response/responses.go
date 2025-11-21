package response

import "github.com/JayBoba/avito-internship/internal/dataflow"

//----- HEALTH

//----- ERRORS
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error *ErrorDetail `json:"error"`
}

func NewErrorResponse(code, message string) *ErrorResponse {
	return &ErrorResponse{
		Error: &ErrorDetail{
			Code:    code,
			Message: message,
		},
	}
}

//-----USER
type SetIsActiveResponse struct {
	User dataflow.UserDTO `json:"user"`
}

type GetReviewResponse struct {
	UserID dataflow.UserDTO        `json:"user"`
	PRs    dataflow.PullRequestDTO //?????????
}

//-----TEAM

//-----PR
