package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源取值范围，最大值最小值可取。 约束：min<=max
type DoubleRange struct {

	// 最大值：999999999.9
	Max *float64 `json:"max,omitempty"`

	// 最小值：0
	Min *float64 `json:"min,omitempty"`
}

func (o DoubleRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DoubleRange struct{}"
	}

	return strings.Join([]string{"DoubleRange", string(data)}, " ")
}
