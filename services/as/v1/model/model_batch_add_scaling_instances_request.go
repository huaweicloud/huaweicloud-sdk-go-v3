package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddScalingInstancesRequest Request Object
type BatchAddScalingInstancesRequest struct {

	// 实例ID。
	ScalingGroupId string `json:"scaling_group_id"`

	Body *BatchAddInstancesOption `json:"body,omitempty"`
}

func (o BatchAddScalingInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddScalingInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchAddScalingInstancesRequest", string(data)}, " ")
}
