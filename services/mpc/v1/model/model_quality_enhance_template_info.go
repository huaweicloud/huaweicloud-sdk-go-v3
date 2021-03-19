package model

import (
	"encoding/json"

	"strings"
)

type QualityEnhanceTemplateInfo struct {
	// 模板ID。

	TemplateId *int32 `json:"template_id,omitempty"`

	Template *QualityEnhanceTemplate `json:"template,omitempty"`

	Error *XCodeError `json:"error,omitempty"`
}

func (o QualityEnhanceTemplateInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QualityEnhanceTemplateInfo struct{}"
	}

	return strings.Join([]string{"QualityEnhanceTemplateInfo", string(data)}, " ")
}
