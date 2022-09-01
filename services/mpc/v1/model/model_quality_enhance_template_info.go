package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QualityEnhanceTemplateInfo struct {

	// 模板ID。
	TemplateId *int32 `json:"template_id,omitempty" xml:"template_id"`

	Template *QualityEnhanceTemplate `json:"template,omitempty" xml:"template"`

	Error *XCodeError `json:"error,omitempty" xml:"error"`
}

func (o QualityEnhanceTemplateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualityEnhanceTemplateInfo struct{}"
	}

	return strings.Join([]string{"QualityEnhanceTemplateInfo", string(data)}, " ")
}
