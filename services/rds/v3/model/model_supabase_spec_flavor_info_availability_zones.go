package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupabaseSpecFlavorInfoAvailabilityZones struct {
	Code *string `json:"code,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o SupabaseSpecFlavorInfoAvailabilityZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseSpecFlavorInfoAvailabilityZones struct{}"
	}

	return strings.Join([]string{"SupabaseSpecFlavorInfoAvailabilityZones", string(data)}, " ")
}
