package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerResource struct {

	// 后端elb实例ID。
	ResourceId string `json:"resource_id"`

	// elb所在的az_id。
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`
}

func (o ServerResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerResource struct{}"
	}

	return strings.Join([]string{"ServerResource", string(data)}, " ")
}
