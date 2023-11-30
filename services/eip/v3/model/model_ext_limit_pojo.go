package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtLimitPojo struct {

	// - 最小入云限速
	MinIngressSize *int32 `json:"min_ingress_size,omitempty"`

	// - 最大入云限速
	MaxIngressSize *int32 `json:"max_ingress_size,omitempty"`

	// 95计费保底率
	Ratio95peak *int32 `json:"ratio_95peak,omitempty"`
}

func (o ExtLimitPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtLimitPojo struct{}"
	}

	return strings.Join([]string{"ExtLimitPojo", string(data)}, " ")
}
