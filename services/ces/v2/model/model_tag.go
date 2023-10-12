package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag 标签
type Tag struct {

	// 键；最大长度128个unicode字符。key不能为空。
	Key string `json:"key"`

	// 值列表；每个值最大长度255个unicode字符。如果values为空列表，则表示查询任意value
	Values []string `json:"values"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
