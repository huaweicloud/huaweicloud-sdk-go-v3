package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowTempRequest struct {
	// 事务id

	TemplateId int32 `json:"template_id"`
}

func (o ShowTempRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTempRequest struct{}"
	}

	return strings.Join([]string{"ShowTempRequest", string(data)}, " ")
}
