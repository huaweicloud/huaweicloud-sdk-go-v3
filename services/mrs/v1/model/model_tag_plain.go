package model

import (
	"encoding/json"

	"strings"
)

type TagPlain struct {
	// 键。标签的key值不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。

	Key *string `json:"key,omitempty"`
	// 值。标签的value值不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。

	Value *string `json:"value,omitempty"`
}

func (o TagPlain) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TagPlain struct{}"
	}

	return strings.Join([]string{"TagPlain", string(data)}, " ")
}
