package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InsertHeaderConfig struct {

	// **参数解释**：请求头参数名。  **取值范围**：1-40个字符，字母a-z（不区分大小写）、数字，短划线-和下划线_。
	Key string `json:"key"`

	// **参数解释**：请求头参数类别。  **取值范围**：USER_DEFINED,REFERENCE_HEADER,SYSTEM_DEFINED。
	ValueType string `json:"value_type"`

	// **参数解释**：请求头参数的值。  **取值范围**：1-128个字符，支持ascii码值32<=ch<=127范围内可打印字符，*和英文问号?。不能以空格开头或结尾。
	Value string `json:"value"`
}

func (o InsertHeaderConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsertHeaderConfig struct{}"
	}

	return strings.Join([]string{"InsertHeaderConfig", string(data)}, " ")
}
