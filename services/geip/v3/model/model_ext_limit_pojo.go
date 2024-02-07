package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtLimitPojo struct {

	// 入网全域公网带宽的最小size
	MinIngressSize *int32 `json:"min_ingress_size,omitempty"`

	// 入网全域公网带宽的最大size
	MaxIngressSize *int32 `json:"max_ingress_size,omitempty"`

	// 增强95保底率
	Ratio95peak *int32 `json:"ratio_95peak,omitempty"`
}

func (o ExtLimitPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtLimitPojo struct{}"
	}

	return strings.Join([]string{"ExtLimitPojo", string(data)}, " ")
}
