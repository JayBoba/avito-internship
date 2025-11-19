package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JayBoba/avito-internship/internal/controller/http/v1/response"
	"github.com/JayBoba/avito-internship/internal/entity"
)

/*
package request

1 Create a Http POST request using http.NewRequest method.
2 The first parameter indicates HTTP request type i.e., “POST”
3 Second parameter is URL of the post request.
4 And the third parameter in request data i.e., JSON data.
5 Set HTTP request header Content-Type as application/json.
6 Finally create a client and make the post request using client.Do(request) method.
*/

type TeamRequest struct {
	TeamName string              `json:"team_name"`
	Members  []entity.TeamMember `json:"members"`
}

func TeamAdd(apiBaseURL string, teamReq TeamRequest) (*response.TeamAddResponse, error) {
	url := fmt.Sprintf("%s/team/add", apiBaseURL)

	jsonData, err := json.Marshal(teamReq)
	if err != nil {
		return nil, fmt.Errorf("request decoding troubles: %w", err) //+LOG later
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("couldn't create a request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("couldn't fullfil the request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading rsponse: %w", err)
	}

	var teamAddResponse response.TeamAddResponse
	if err := json.Unmarshal(body, &teamAddResponse); err != nil {
		return nil, fmt.Errorf("couldn't decode: %w", err)
	}

	return &teamAddResponse, nil
}

/*
  /team/get:
    get:
      tags: [Teams]   
      summary: Получить команду с участниками
      parameters:
        - $ref: '#/components/parameters/TeamNameQuery'
*/
