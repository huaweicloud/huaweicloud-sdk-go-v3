package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchUnprotectScalingInstancesRequest struct {

	// 实例ID。
	ScalingGroupId string `json:"scaling_group_id" xml:"scaling_group_id"`

	Body *BatchUnprotectInstancesOption `json:"body,omitempty" xml:"body"`
}

func (o BatchUnprotectScalingInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUnprotectScalingInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchUnprotectScalingInstancesRequest", string(data)}, " ")
}
