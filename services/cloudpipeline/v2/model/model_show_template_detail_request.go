package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTemplateDetailRequest struct {
	// 语言类型 中文:zh-cn 英文:en-us，默认en-us

	XLanguage *string `json:"X-Language,omitempty"`
	// 模板ID

	TemplateId string `json:"template_id"`
	// 模板类型

	TemplateType string `json:"template_type"`
	// 接口调用方

	Source *string `json:"source,omitempty"`
}

func (o ShowTemplateDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateDetailRequest", string(data)}, " ")
}
