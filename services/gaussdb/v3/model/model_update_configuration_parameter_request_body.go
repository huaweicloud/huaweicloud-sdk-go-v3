package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateConfigurationParameterRequestBody struct {

	// 参数模板名称。  取值范围：长度1到64个字符之间，区分大小写字母，可包含字母、数字、中划线、下划线或句点，不能包含其他特殊字符。参数模板描述，参数名和参数值映射关系三项不能同时为空。
	Name *string `json:"name,omitempty"`

	// 参数模板描述。默认为空。取值范围：长度不超过256个字符，且不能包含回车和特殊字符 ! < \" = ' > &
	Description *string `json:"description,omitempty"`

	// 参数名和参数值映射关系。用户可以基于默认参数模板的参数，自定义的参数值。不传入该参数，则保持原参数信息。
	ParameterValues map[string]string `json:"parameter_values,omitempty"`
}

func (o UpdateConfigurationParameterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationParameterRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationParameterRequestBody", string(data)}, " ")
}
