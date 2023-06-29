package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZoneRequest Request Object
type ListAvailabilityZoneRequest struct {
}

func (o ListAvailabilityZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZoneRequest struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZoneRequest", string(data)}, " ")
}
