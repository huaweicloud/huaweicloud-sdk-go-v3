package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateGaussMySqlProxyResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGaussMySqlProxyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlProxyResponse struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlProxyResponse", string(data)}, " ")
}
