package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackgroundTaskResponse Response Object
type ListBackgroundTaskResponse struct {

	// 任务个数
	TaskCount *string `json:"task_count,omitempty"`

	// 任务详情数组
	Tasks          *[]SingleBackgroundTask `json:"tasks,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListBackgroundTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackgroundTaskResponse struct{}"
	}

	return strings.Join([]string{"ListBackgroundTaskResponse", string(data)}, " ")
}
