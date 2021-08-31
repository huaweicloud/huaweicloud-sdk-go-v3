package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateTempRequest struct {
	// 事务id

	TemplateId int32 `json:"template_id"`

	Body *UpdateTempRequestBody `json:"body,omitempty"`
}

func (o UpdateTempRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTempRequest struct{}"
	}

	return strings.Join([]string{"UpdateTempRequest", string(data)}, " ")
}
