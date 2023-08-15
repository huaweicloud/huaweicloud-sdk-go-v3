package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailabilityZone 可用区。
type AvailabilityZone struct {

	// 可用区唯一编码。
	Code string `json:"code"`

	// 可用区状态。  取值：ACTIVE。
	State string `json:"state"`

	// 未售罄的LB规格类别。  取值：L4 表示网络型LB未售罄；L7 表示应用型LB未售罄。
	Protocol []string `json:"protocol"`

	// 可用区组，如：center
	PublicBorderGroup string `json:"public_border_group"`

	// 范围编码，0表示center，21表示homezone
	Category int32 `json:"category"`
}

func (o AvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZone struct{}"
	}

	return strings.Join([]string{"AvailabilityZone", string(data)}, " ")
}
