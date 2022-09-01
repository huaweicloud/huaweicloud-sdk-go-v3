package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// tag字段说明
type Tag struct {

	// 标签的键。  最大长度127个字符。 key不能为空。
	Key string `json:"key" xml:"key"`

	// 值列表。每个值最大长度255个字符，如果values为空列表，则表示匹配任意值value。value之间为或的关系。
	Values []string `json:"values" xml:"values"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
