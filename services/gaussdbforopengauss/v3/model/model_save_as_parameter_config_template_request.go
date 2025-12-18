package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveAsParameterConfigTemplateRequest Request Object
type SaveAsParameterConfigTemplateRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 参数模板ID。
	ConfigId string `json:"config_id"`

	Body *SaveAsParameterConfigTemplateRequestBody `json:"body,omitempty"`
}

func (o SaveAsParameterConfigTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveAsParameterConfigTemplateRequest struct{}"
	}

	return strings.Join([]string{"SaveAsParameterConfigTemplateRequest", string(data)}, " ")
}
