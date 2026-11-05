package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupabaseSpecFlavorInfo 规格信息，含可售AZ和版本。
type SupabaseSpecFlavorInfo struct {
	Id *string `json:"id,omitempty"`

	Mode *string `json:"mode,omitempty"`

	Version *string `json:"version,omitempty"`

	Vcpus *int32 `json:"vcpus,omitempty"`

	Ram *int32 `json:"ram,omitempty"`

	SpecCode *string `json:"spec_code,omitempty"`

	GroupType *string `json:"group_type,omitempty"`

	AvailabilityZones *[]SupabaseSpecFlavorInfoAvailabilityZones `json:"availability_zones,omitempty"`
}

func (o SupabaseSpecFlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseSpecFlavorInfo struct{}"
	}

	return strings.Join([]string{"SupabaseSpecFlavorInfo", string(data)}, " ")
}
