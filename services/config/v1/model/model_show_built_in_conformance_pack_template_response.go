package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuiltInConformancePackTemplateResponse Response Object
type ShowBuiltInConformancePackTemplateResponse struct {

	// 预定义合规包模板ID。
	Id *string `json:"id,omitempty"`

	// 预定义合规包模板名称。
	TemplateKey *string `json:"template_key,omitempty"`

	// 预定义合规包模板描述。
	Description *string `json:"description,omitempty"`

	// 预定义合规包模板内容。
	TemplateBody *string `json:"template_body,omitempty"`

	// 预定义合规包模板参数。
	Parameters     map[string]TemplateParameterDefinition `json:"parameters,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ShowBuiltInConformancePackTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuiltInConformancePackTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowBuiltInConformancePackTemplateResponse", string(data)}, " ")
}
