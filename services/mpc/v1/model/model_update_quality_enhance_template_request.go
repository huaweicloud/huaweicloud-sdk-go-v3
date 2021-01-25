package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateQualityEnhanceTemplateRequest struct {
	Body *UpdateQualityEnhanceTemplateReq `json:"body,omitempty"`
}

func (o UpdateQualityEnhanceTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateQualityEnhanceTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateQualityEnhanceTemplateRequest", string(data)}, " ")
}
