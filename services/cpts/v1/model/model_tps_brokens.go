package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TpsBrokens struct {

	// 平均响应时间
	Average *[]float64 `json:"average,omitempty" xml:"average"`

	// tps
	Tps *[]float64 `json:"tps,omitempty" xml:"tps"`
}

func (o TpsBrokens) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TpsBrokens struct{}"
	}

	return strings.Join([]string{"TpsBrokens", string(data)}, " ")
}
