package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 自定义标签键值对。
type TagsDto struct {

	// 键。最大长度127个unicode字符。 key不能为空。
	Key string `json:"key"`

	// 值列表。每个值最大长度255个unicode字符。
	Values []string `json:"values"`
}

func (o TagsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsDto struct{}"
	}

	return strings.Join([]string{"TagsDto", string(data)}, " ")
}
