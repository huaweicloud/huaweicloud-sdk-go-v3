package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConfigurationTemplateRequestBody struct {

	// 参数模板名称，不可与已有参数模板名称重复。 取值范围：长度1到64位之间，区分大小写字母，可包含字母、数字、中划线、下划线或句点，不能包含其他特殊字符。
	Name string `json:"name"`

	// 参数模板描述，默认为空。 取值范围：长度不超过256，不能包含回车<>!&等特殊字符。
	Description *string `json:"description,omitempty"`

	// 参数名和参数值映射关系。用户可以基于默认参数模板的参数，自定义参数值。
	ParameterValues map[string]string `json:"parameter_values,omitempty"`

	Datastore *DatastoreOption `json:"datastore"`
}

func (o CreateConfigurationTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConfigurationTemplateRequestBody", string(data)}, " ")
}
