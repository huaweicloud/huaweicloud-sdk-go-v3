package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteGaussMySqlReadonlyNodeResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGaussMySqlReadonlyNodeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlReadonlyNodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlReadonlyNodeResponse", string(data)}, " ")
}
