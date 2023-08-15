package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZonesRequest Request Object
type ListAvailabilityZonesRequest struct {
}

func (o ListAvailabilityZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZonesRequest struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZonesRequest", string(data)}, " ")
}
