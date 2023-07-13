package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackgroundTaskRequest Request Object
type ListBackgroundTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 查询开始时间，时间为UTC时间。格式：yyyyMMddHHmmss，如：20200609160000。
	StartTime *string `json:"start_time,omitempty"`

	// 查询结束时间，时间为UTC时间。格式：yyyyMMddHHmmss，如：20200617155959。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListBackgroundTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackgroundTaskRequest struct{}"
	}

	return strings.Join([]string{"ListBackgroundTaskRequest", string(data)}, " ")
}
