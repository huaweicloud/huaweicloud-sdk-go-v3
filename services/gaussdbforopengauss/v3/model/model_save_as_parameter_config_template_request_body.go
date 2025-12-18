package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SaveAsParameterConfigTemplateRequestBody struct {

	// 保存参数组模板的名称。
	Name *string `json:"name,omitempty"`

	// 保存参数组模板的描述。
	Description *string `json:"description,omitempty"`
}

func (o SaveAsParameterConfigTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveAsParameterConfigTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"SaveAsParameterConfigTemplateRequestBody", string(data)}, " ")
}
