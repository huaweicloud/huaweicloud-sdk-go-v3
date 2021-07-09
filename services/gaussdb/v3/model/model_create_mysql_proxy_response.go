package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateMysqlProxyResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMysqlProxyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMysqlProxyResponse struct{}"
	}

	return strings.Join([]string{"CreateMysqlProxyResponse", string(data)}, " ")
}
