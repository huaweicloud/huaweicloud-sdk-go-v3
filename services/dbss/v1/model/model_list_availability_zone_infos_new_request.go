package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZoneInfosNewRequest Request Object
type ListAvailabilityZoneInfosNewRequest struct {
}

func (o ListAvailabilityZoneInfosNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZoneInfosNewRequest struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZoneInfosNewRequest", string(data)}, " ")
}
