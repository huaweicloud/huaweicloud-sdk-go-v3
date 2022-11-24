package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 属性权重
type NumericalAttr struct {

	// 特征名。
	Name string `json:"name"`

	// 权重。
	Weight float32 `json:"weight"`
}

func (o NumericalAttr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NumericalAttr struct{}"
	}

	return strings.Join([]string{"NumericalAttr", string(data)}, " ")
}
