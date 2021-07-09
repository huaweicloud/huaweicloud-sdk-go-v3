package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ExpandMysqlProxyResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandMysqlProxyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExpandMysqlProxyResponse struct{}"
	}

	return strings.Join([]string{"ExpandMysqlProxyResponse", string(data)}, " ")
}
