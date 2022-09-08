package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateConfigurationParameterRequestBody struct {

	// 参数模板名称。最长64个字符，只允许大写字母、小写字母、数字和特殊字符中划线、下划线和点。不传该参数时不修改参数模板名称。（参数模板名称，参数模板描述，参数名和参数值映射关系三项不能同时为空）
	Name *string `json:"name,omitempty"`

	// 参数模板描述。最长256位，不支持!<>=&\"'特殊字符。不传该参数时不修改参数模板描述。
	Description *string `json:"description,omitempty"`

	// 参数名和参数值映射关系。用户可以基于默认参数模板的参数，自定义的参数值。
	ParameterValues map[string]string `json:"parameter_values,omitempty"`
}

func (o UpdateConfigurationParameterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationParameterRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationParameterRequestBody", string(data)}, " ")
}
