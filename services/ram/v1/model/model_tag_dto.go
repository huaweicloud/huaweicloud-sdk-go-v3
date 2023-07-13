package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagDto 自定义标签键值对。
type TagDto struct {

	// 标签\"键\"的标识符或名称。您可以将标签的值设置为空字符串，但不能设置为null。
	Key string `json:"key"`

	// 值列表。每个值最大长度255个字符。
	Values []string `json:"values"`
}

func (o TagDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagDto struct{}"
	}

	return strings.Join([]string{"TagDto", string(data)}, " ")
}
