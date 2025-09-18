package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SpecificationCodeInfo struct {

	// 带宽包实例的规格编码。
	SpecCode string `json:"spec_code"`

	BillingMode *BillingModeEnum `json:"billing_mode"`

	// 最大带宽。
	MaxBandwidth *int64 `json:"max_bandwidth,omitempty"`

	// 最小带宽。
	MimBandwidth *int64 `json:"mim_bandwidth,omitempty"`
}

func (o SpecificationCodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecificationCodeInfo struct{}"
	}

	return strings.Join([]string{"SpecificationCodeInfo", string(data)}, " ")
}
