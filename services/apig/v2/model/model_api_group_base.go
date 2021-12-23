package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiGroupBase struct {
	// API分组的名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头，3-255个字符。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// API分组描述。 > 中文字符必须为UTF-8或者unicode编码。

	Remark *string `json:"remark,omitempty"`
}

func (o ApiGroupBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiGroupBase struct{}"
	}

	return strings.Join([]string{"ApiGroupBase", string(data)}, " ")
}
