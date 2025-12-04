package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksRequest Request Object
type ListTasksRequest struct {

	// 开始时间。
	StartTime float32 `json:"start_time"`

	// 开始时间。
	EndTime float32 `json:"end_time"`
}

func (o ListTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequest struct{}"
	}

	return strings.Join([]string{"ListTasksRequest", string(data)}, " ")
}
