package model

import (
	"encoding/json"

	"strings"
)

type OpIdRes struct {
	// 操作记录id

	OperationId *string `json:"operation_id,omitempty"`
}

func (o OpIdRes) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OpIdRes struct{}"
	}

	return strings.Join([]string{"OpIdRes", string(data)}, " ")
}
