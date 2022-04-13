package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Tag struct {
	// 键。 key最大长度为36个字符。 key不能为空字符串。 key前后空格会被丢弃。 key不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。 key只能由中文，字母，数字，“-”，“_”组成。

	Key string `json:"key"`
	// 值。 添加标签时value值必选，删除标签时value值可选。 value最大长度为43个字符。 value可以为空字符串。 value前后的空格会被丢弃。 value不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。 value只能由中文，字母，数字，“-”，“_”，“.”组成。

	Value *string `json:"value,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
