package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tag struct {
	// 键。 - 最大长度36个unicode字符，不能为空字符串。  - 标签的key值不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。  - 同一资源的key值不能重复。

	Key string `json:"key"`
	// 值。 - 最大长度43个unicode字符，可以为空字符串。  - 标签的value值不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。

	Value string `json:"value"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
