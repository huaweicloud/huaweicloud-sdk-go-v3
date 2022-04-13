package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobDetailResponse struct {
	// 部署任务数量。

	TaskCount *int32 `json:"task_count,omitempty"`

	Job *JobInfo `json:"job,omitempty"`
	// 部署任务列表。

	Tasks          *[]TaskInfo `json:"tasks,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowJobDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowJobDetailResponse", string(data)}, " ")
}
