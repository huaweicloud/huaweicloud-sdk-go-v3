package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Untag 自定义键值对。
type Untag struct {

	// 标签\"键\"的标识符或名称。
	Key string `json:"key"`

	// 标签\"键\"对应的\"值\"。您可以将标签的值设置为空字符串，但不能设置为null。
	Value *string `json:"value,omitempty"`
}

func (o Untag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Untag struct{}"
	}

	return strings.Join([]string{"Untag", string(data)}, " ")
}
