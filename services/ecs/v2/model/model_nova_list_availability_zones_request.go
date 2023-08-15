package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaListAvailabilityZonesRequest Request Object
type NovaListAvailabilityZonesRequest struct {
}

func (o NovaListAvailabilityZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListAvailabilityZonesRequest struct{}"
	}

	return strings.Join([]string{"NovaListAvailabilityZonesRequest", string(data)}, " ")
}
