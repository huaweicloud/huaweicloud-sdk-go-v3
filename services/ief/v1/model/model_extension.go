package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Extension 批量处理对象额外属性
type Extension struct {

	// 属性名，可填：node_name
	Key *string `json:"key,omitempty"`

	// 属性值
	Value *string `json:"value,omitempty"`
}

func (o Extension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Extension struct{}"
	}

	return strings.Join([]string{"Extension", string(data)}, " ")
}
