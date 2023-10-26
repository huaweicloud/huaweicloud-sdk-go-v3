package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCenterTaskResponse Response Object
type ListCenterTaskResponse struct {

	// 任务个数
	TaskCount *string `json:"task_count,omitempty"`

	// 中心任务详情数组
	Tasks          *[]ListCenterTasksResp `json:"tasks,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListCenterTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCenterTaskResponse struct{}"
	}

	return strings.Join([]string{"ListCenterTaskResponse", string(data)}, " ")
}
