package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAvailabilityZoneRequest struct {

	// 企业路由器ID
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ListAvailabilityZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZoneRequest struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZoneRequest", string(data)}, " ")
}
