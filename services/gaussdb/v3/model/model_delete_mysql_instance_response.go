package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteMysqlInstanceResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMysqlInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMysqlInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteMysqlInstanceResponse", string(data)}, " ")
}
