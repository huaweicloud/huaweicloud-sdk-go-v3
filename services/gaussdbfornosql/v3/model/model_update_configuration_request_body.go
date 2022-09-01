package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateConfigurationRequestBody struct {

	// 参数模板名称。最长64个字符，只允许大写字母、小写字母、数字和特殊字符中划线、下划线和点。
	Name *string `json:"name,omitempty" xml:"name"`

	// 参数模板描述。最长256个字符，不支持!<>=&\"'特殊字符。默认为空。
	Description *string `json:"description,omitempty" xml:"description"`

	// 参数值对象，用户基于默认参数模板自定义的参数值。为空时不修改参数值。
	Values map[string]string `json:"values,omitempty" xml:"values"`
}

func (o UpdateConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationRequestBody", string(data)}, " ")
}
