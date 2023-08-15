package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CopyConfigurationRequestBody struct {

	// 复制后的参数模板名称。最长64个字符，只允许大写字母、小写字母、数字、和“-_.”特殊字符。
	Name string `json:"name"`

	// 参数模板描述。最长256个字符，不支持>!<\"&'=特殊字符。默认为空。
	Description *string `json:"description,omitempty"`
}

func (o CopyConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CopyConfigurationRequestBody", string(data)}, " ")
}
