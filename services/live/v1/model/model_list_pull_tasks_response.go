package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPullTasksResponse Response Object
type ListPullTasksResponse struct {

	// 直播拉流任务列表
	TaskInfos *[]LivePullTaskInfo `json:"task_infos,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	// 任务条目数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPullTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPullTasksResponse struct{}"
	}

	return strings.Join([]string{"ListPullTasksResponse", string(data)}, " ")
}
