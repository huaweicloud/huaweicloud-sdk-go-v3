package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteQualityEnhanceTemplateRequest struct {
	// 模板ID

	TemplateId int32 `json:"template_id"`
}

func (o DeleteQualityEnhanceTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteQualityEnhanceTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteQualityEnhanceTemplateRequest", string(data)}, " ")
}
