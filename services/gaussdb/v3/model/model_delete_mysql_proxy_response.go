package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteMysqlProxyResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMysqlProxyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMysqlProxyResponse struct{}"
	}

	return strings.Join([]string{"DeleteMysqlProxyResponse", string(data)}, " ")
}
