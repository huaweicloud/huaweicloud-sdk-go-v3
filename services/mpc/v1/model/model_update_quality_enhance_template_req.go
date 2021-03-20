package model

import (
	"encoding/json"

	"strings"
)

type UpdateQualityEnhanceTemplateReq struct {
	// 模板ID。

	TemplateId *int32 `json:"template_id,omitempty"`

	Template *QualityEnhanceTemplate `json:"template,omitempty"`
}

func (o UpdateQualityEnhanceTemplateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateQualityEnhanceTemplateReq struct{}"
	}

	return strings.Join([]string{"UpdateQualityEnhanceTemplateReq", string(data)}, " ")
}
