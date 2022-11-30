package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComputeFlavor struct {

	// cpu核数。
	Vcpus *string `json:"vcpus,omitempty"`

	// 内存大小，单位为GB。
	Ram *string `json:"ram,omitempty"`

	// 规格码。
	SpecCode *string `json:"spec_code,omitempty"`

	// 可用区状态。
	AzStatus map[string]string `json:"az_status,omitempty"`

	// Region状态。
	RegionStatus *string `json:"region_status,omitempty"`
}

func (o ComputeFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeFlavor struct{}"
	}

	return strings.Join([]string{"ComputeFlavor", string(data)}, " ")
}
