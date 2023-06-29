package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag tag字段数据结构。
type Tag struct {

	// 键。  最大长度127个unicode字符。  key不能为空。
	Key string `json:"key"`

	// 值列表。  最多10个value。  value不允许重复。  每个值最大长度255个unicode字符。  如果value缺失则表示any_value。  value之间为“或”的关系。
	Values *[]string `json:"values,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
