package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数值范围参数结构体。
type RangeParam struct {

	// 数值下界，默认包含该下界。
	From *float64 `json:"from,omitempty"`

	// 数值上界，默认包含该上界。
	To *float64 `json:"to,omitempty"`
}

func (o RangeParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RangeParam struct{}"
	}

	return strings.Join([]string{"RangeParam", string(data)}, " ")
}
