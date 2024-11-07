package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthPackageLineSpecCode 带宽包线路产品
type BandwidthPackageLineSpecCode struct {

	// 带宽包等级
	Level *string `json:"level,omitempty"`

	// 实例名字。
	NameCn *string `json:"name_cn,omitempty"`

	// 实例名字。
	NameEn *string `json:"name_en,omitempty"`

	// 带宽包实例的规格编码。
	SpecCode *string `json:"spec_code,omitempty"`

	// 最大带宽。
	MaxBandwidth *int32 `json:"max_bandwidth,omitempty"`

	// 最小带宽。
	MinBandwidth *int32 `json:"min_bandwidth,omitempty"`

	// 支持的计费模式。
	SupportBillingModes *[]BillingModeEnum `json:"support_billing_modes,omitempty"`
}

func (o BandwidthPackageLineSpecCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthPackageLineSpecCode struct{}"
	}

	return strings.Join([]string{"BandwidthPackageLineSpecCode", string(data)}, " ")
}
