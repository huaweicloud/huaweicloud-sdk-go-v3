package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateQualityEnhanceTemplateReq struct {

	// 模板ID。
	TemplateId *int32 `json:"template_id,omitempty" xml:"template_id"`

	Template *QualityEnhanceTemplate `json:"template,omitempty" xml:"template"`
}

func (o UpdateQualityEnhanceTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQualityEnhanceTemplateReq struct{}"
	}

	return strings.Join([]string{"UpdateQualityEnhanceTemplateReq", string(data)}, " ")
}
