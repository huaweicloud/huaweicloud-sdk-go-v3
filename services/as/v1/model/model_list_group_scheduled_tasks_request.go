package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupScheduledTasksRequest Request Object
type ListGroupScheduledTasksRequest struct {

	// 伸缩组ID
	ScalingGroupId string `json:"scaling_group_id"`

	// 查询的记录条数
	Limit *int32 `json:"limit,omitempty"`

	// 查询的分页marker
	Marker *string `json:"marker,omitempty"`
}

func (o ListGroupScheduledTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupScheduledTasksRequest struct{}"
	}

	return strings.Join([]string{"ListGroupScheduledTasksRequest", string(data)}, " ")
}
