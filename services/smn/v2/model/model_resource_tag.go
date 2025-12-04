package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTag 资源标签结构体。
type ResourceTag struct {

	// 键。  - 最大长度127个unicode字符。  - key不能为空。
	Key string `json:"key"`

	// 值。  - 每个值最大长度255个unicode字符。
	Value string `json:"value"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
