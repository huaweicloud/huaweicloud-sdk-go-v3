package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteGaussMySqlProxyResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGaussMySqlProxyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlProxyResponse struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlProxyResponse", string(data)}, " ")
}
