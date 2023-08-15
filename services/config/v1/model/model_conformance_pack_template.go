package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConformancePackTemplate 预定义合规规则包模板详情。
type ConformancePackTemplate struct {

	// 预定义合规包模板ID。
	Id *string `json:"id,omitempty"`

	// 预定义合规包模板名称。
	TemplateKey *string `json:"template_key,omitempty"`

	// 预定义合规包模板描述。
	Description *string `json:"description,omitempty"`

	// 预定义合规包模板内容。
	TemplateBody *string `json:"template_body,omitempty"`

	// 预定义合规包模板参数。
	Parameters map[string]TemplateParameterDefinition `json:"parameters,omitempty"`
}

func (o ConformancePackTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConformancePackTemplate struct{}"
	}

	return strings.Join([]string{"ConformancePackTemplate", string(data)}, " ")
}
