package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagWithMultiValue struct {
	// 键。标签的key值不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。

	Key string `json:"key"`
	// 值。标签的value值不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。

	Values *[]string `json:"values,omitempty"`
}

func (o TagWithMultiValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagWithMultiValue struct{}"
	}

	return strings.Join([]string{"TagWithMultiValue", string(data)}, " ")
}
