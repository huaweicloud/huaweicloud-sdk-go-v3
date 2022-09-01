package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签列表，没有标签默认为空数组。
type TagList struct {

	// 键。最大长度36个unicode字符。 key不 能为空。不能包含“=”、“*”、 “<”、“>”、“\\”、“,”、“|”和 “/”，且首尾字符不能为空格。
	Key *string `json:"key,omitempty" xml:"key"`

	// 值。每个值最大长度43个unicode字 符，可以为空字符串。 不能包含 “=”、“*”、“<”、“>”、“\\”、 “,”、“|”和“/”，且首尾字符不能 为空格。
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o TagList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagList struct{}"
	}

	return strings.Join([]string{"TagList", string(data)}, " ")
}
