package response

import "github.com/JayBoba/avito-internship/internal/entity"

type TeamAddResponse struct {
	Team *entity.Team `json:"team"`
}

func NewTeamAddResponse(team *entity.Team) *TeamAddResponse {
	return &TeamAddResponse{
		Team: team,
	}
}

type ErrorDetail struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
}

type ErrorResponse struct {
	Error *ErrorDetail `json:"error"`
}

func NewErrorResponse(code, message string) *ErrorResponse {
	return &ErrorResponse{
		Error: &ErrorDetail{
			Code:        code,
			Message:     message,
		},
	}
}

/* get
   responses:
     '200':
       description: Объект команды
       content:
         application/json:
           schema:
             $ref: '#/components/schemas/Team'
           example:
             team_name: backend
             members:
               - user_id: u1
                 username: Alice
                 is_active: true
               - user_id: u2
                 username: Bob
                 is_active: true
     '404':
       description: Команда не найдена
       content:
         application/json:
           schema: { $ref: '#/components/schemas/ErrorResponse' }
*/
