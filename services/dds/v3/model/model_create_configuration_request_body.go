package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConfigurationRequestBody struct {

	// 参数模板名称。 取值范围：长度1到64位之间，区分大小写字母，可包含字母、数字、中划线、下划线或句点，不能包含其他特殊字符。
	Name string `json:"name"`

	// 参数模板描述。 取值范围：长度不超过256位，且不能包含回车和>!<\"&'=特殊字符。默认为空
	Description string `json:"description"`

	// 参数名和参数值映射关系。用户可以基于默认参数模板的参数，自定义的参数值。
	ParameterValues map[string]string `json:"parameter_values"`

	Datastore *DatastoreResult `json:"datastore"`
}

func (o CreateConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConfigurationRequestBody", string(data)}, " ")
}
