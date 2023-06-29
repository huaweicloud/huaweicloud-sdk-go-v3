package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DoubleRange 资源取值范围，最大值最小值可取。 约束：min<=max
type DoubleRange struct {

	// 最大值，-1表示无限制
	Max *float64 `json:"max,omitempty"`

	// 最小值，-1表示无限制
	Min *float64 `json:"min,omitempty"`
}

func (o DoubleRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DoubleRange struct{}"
	}

	return strings.Join([]string{"DoubleRange", string(data)}, " ")
}
