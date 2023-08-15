package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZoneInfosRequest Request Object
type ListAvailabilityZoneInfosRequest struct {
}

func (o ListAvailabilityZoneInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZoneInfosRequest struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZoneInfosRequest", string(data)}, " ")
}
