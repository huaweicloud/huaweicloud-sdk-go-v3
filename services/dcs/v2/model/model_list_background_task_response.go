package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBackgroundTaskResponse struct {

	// 任务个数
	TaskCount *string `json:"task_count,omitempty"`

	// 任务详情数组
	Tasks *[]SingleBackgroundTask `json:"tasks,omitempty"`

	// 任务结束时间，格式为2020-06-17T07:38:42.503Z
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 任务启动时间，格式为2020-06-17T07:38:42.503Z
	CreatedAt *string `json:"created_at,omitempty"`

	// 任务状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBackgroundTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackgroundTaskResponse struct{}"
	}

	return strings.Join([]string{"ListBackgroundTaskResponse", string(data)}, " ")
}
