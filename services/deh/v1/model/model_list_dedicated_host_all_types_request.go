package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDedicatedHostAllTypesRequest Request Object
type ListDedicatedHostAllTypesRequest struct {
	Flavor *string `json:"flavor,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	HostType *string `json:"host_type,omitempty"`
}

func (o ListDedicatedHostAllTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostAllTypesRequest struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostAllTypesRequest", string(data)}, " ")
}
