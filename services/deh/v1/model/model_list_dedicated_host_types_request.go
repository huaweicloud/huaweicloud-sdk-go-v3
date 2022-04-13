package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDedicatedHostTypesRequest struct {
	// AZ。

	AvailabilityZone string `json:"availability_zone"`
}

func (o ListDedicatedHostTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostTypesRequest struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostTypesRequest", string(data)}, " ")
}
