package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailabilityZone 可用区详情。
type AvailabilityZone struct {

	// 可用区唯一编码。
	Code string `json:"code"`

	// 可用区名称。
	Name string `json:"name"`

	// 可用区状态。 - available：可用。 - unavailable：不可用。
	Status string `json:"status"`

	// 可用区组，如：center。
	PublicBorderGroup string `json:"public_border_group"`
}

func (o AvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZone struct{}"
	}

	return strings.Join([]string{"AvailabilityZone", string(data)}, " ")
}
