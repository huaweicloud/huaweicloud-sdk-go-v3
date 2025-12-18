package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyParameterConfigTemplateRequestBody struct {

	// 参数值对象Map<String,String>，用户基于默认参数模板自定义的参数值。
	Values map[string]string `json:"values,omitempty"`

	// 参数组模板描述。
	Description *string `json:"description,omitempty"`
}

func (o ModifyParameterConfigTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyParameterConfigTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyParameterConfigTemplateRequestBody", string(data)}, " ")
}
