package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag 标签列表。
type Tag struct {

	// 标签类型: - user - system
	Type string `json:"type"`

	// 标签键。最大长度36个unicode字符，key不能为空。  字符集：0-9，A-Z，a-z，“_”，“-”，中文。
	Key string `json:"key"`

	//  标签值列表。每个标签值最大长度43个unicode字符，可以为空字符串。  字符集：0-9，A-Z，a-z，“_”，“.”，“-”，中文。
	Values []string `json:"values"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
