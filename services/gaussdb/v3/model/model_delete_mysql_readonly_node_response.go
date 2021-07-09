package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteMysqlReadonlyNodeResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMysqlReadonlyNodeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMysqlReadonlyNodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteMysqlReadonlyNodeResponse", string(data)}, " ")
}
