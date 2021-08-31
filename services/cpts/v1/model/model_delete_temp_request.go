package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteTempRequest struct {
	// 事务id

	TemplateId int32 `json:"template_id"`
}

func (o DeleteTempRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTempRequest struct{}"
	}

	return strings.Join([]string{"DeleteTempRequest", string(data)}, " ")
}
