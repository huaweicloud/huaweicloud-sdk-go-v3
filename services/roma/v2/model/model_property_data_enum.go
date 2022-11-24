package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据类型为boolean时的布尔值
type PropertyDataEnum struct {

	// boolean枚举值为0时对应的值
	BoolFalse *string `json:"bool_false,omitempty"`

	// boolean枚举值为1时对应的值
	BoolTrue *string `json:"bool_true,omitempty"`
}

func (o PropertyDataEnum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropertyDataEnum struct{}"
	}

	return strings.Join([]string{"PropertyDataEnum", string(data)}, " ")
}
