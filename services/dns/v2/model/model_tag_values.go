package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagValues struct {
	// 键。最大长度36个unicode字符。 key不能为空。不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。

	Key *string `json:"key,omitempty"`
	// 值列表。每个值最大长度43个unicode字符，可以为空字符串。 不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。

	Values *[]string `json:"values,omitempty"`
}

func (o TagValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagValues struct{}"
	}

	return strings.Join([]string{"TagValues", string(data)}, " ")
}
