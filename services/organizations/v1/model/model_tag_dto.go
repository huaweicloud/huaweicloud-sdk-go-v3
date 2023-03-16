package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 自定义标签键值对。
type TagDto struct {

	// 标签键的密钥标识符或名称。
	Key string `json:"key"`

	// 与标签键关联的字符串值。您可以将标签的值设置为空字符串，但不能将标签的值设置为NULL。
	Value string `json:"value"`
}

func (o TagDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagDto struct{}"
	}

	return strings.Join([]string{"TagDto", string(data)}, " ")
}
