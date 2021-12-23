package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TpsBrokens struct {
	// average

	Average *[]int32 `json:"average,omitempty"`
	// tps

	Tps *[]int32 `json:"tps,omitempty"`
}

func (o TpsBrokens) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TpsBrokens struct{}"
	}

	return strings.Join([]string{"TpsBrokens", string(data)}, " ")
}
