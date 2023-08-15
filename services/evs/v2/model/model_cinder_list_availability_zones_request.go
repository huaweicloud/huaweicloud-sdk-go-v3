package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CinderListAvailabilityZonesRequest Request Object
type CinderListAvailabilityZonesRequest struct {
}

func (o CinderListAvailabilityZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListAvailabilityZonesRequest struct{}"
	}

	return strings.Join([]string{"CinderListAvailabilityZonesRequest", string(data)}, " ")
}
