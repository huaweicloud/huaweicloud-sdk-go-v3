package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExtendedAvailabilityZonesRequest Request Object
type ListExtendedAvailabilityZonesRequest struct {
}

func (o ListExtendedAvailabilityZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExtendedAvailabilityZonesRequest struct{}"
	}

	return strings.Join([]string{"ListExtendedAvailabilityZonesRequest", string(data)}, " ")
}
