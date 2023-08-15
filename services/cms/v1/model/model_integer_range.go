package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IntegerRange 资源取值范围，最大值最小值可取。 约束：min<=max
type IntegerRange struct {

	// 最大值，-1表示无限制
	Max *int32 `json:"max,omitempty"`

	// 最小值，-1表示无限制
	Min *int32 `json:"min,omitempty"`
}

func (o IntegerRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntegerRange struct{}"
	}

	return strings.Join([]string{"IntegerRange", string(data)}, " ")
}
