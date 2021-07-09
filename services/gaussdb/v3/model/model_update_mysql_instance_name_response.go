package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateMysqlInstanceNameResponse struct {
	// 修改实例名称的任务id

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMysqlInstanceNameResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateMysqlInstanceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateMysqlInstanceNameResponse", string(data)}, " ")
}
