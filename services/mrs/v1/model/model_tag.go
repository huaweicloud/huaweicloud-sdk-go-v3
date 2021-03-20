package model

import (
	"encoding/json"

	"strings"
)

type Tag struct {
	// 键。标签的key值不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。

	Key string `json:"key"`
	// 值。标签的value值不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。

	Value string `json:"value"`
}

func (o Tag) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
