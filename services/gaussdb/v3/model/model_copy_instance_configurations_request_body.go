package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CopyInstanceConfigurationsRequestBody struct {

	// 复制后的参数模板名称。  取值范围：长度1到64个字符之间，区分大小写字母，可包含字母、数字、中划线、下划线或句点，不能包含其他特殊字符。
	Name string `json:"name"`

	// 复制后的参数模板描述。默认为空。  取值范围：描述不能超过256个字符，且不能包含回车和特殊字符 ! < \" = ' > &。
	Description *string `json:"description,omitempty"`
}

func (o CopyInstanceConfigurationsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyInstanceConfigurationsRequestBody struct{}"
	}

	return strings.Join([]string{"CopyInstanceConfigurationsRequestBody", string(data)}, " ")
}
