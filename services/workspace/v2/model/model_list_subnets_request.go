package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubnetsRequest Request Object
type ListSubnetsRequest struct {

	// 虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 可用区id。
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`
}

func (o ListSubnetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetsRequest struct{}"
	}

	return strings.Join([]string{"ListSubnetsRequest", string(data)}, " ")
}
