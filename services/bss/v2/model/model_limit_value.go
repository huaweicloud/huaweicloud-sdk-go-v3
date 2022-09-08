package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LimitValue struct {

	// 属性值1。
	Value1 *string `json:"value1,omitempty"`

	// 属性值2。
	Value2 *string `json:"value2,omitempty"`
}

func (o LimitValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LimitValue struct{}"
	}

	return strings.Join([]string{"LimitValue", string(data)}, " ")
}
