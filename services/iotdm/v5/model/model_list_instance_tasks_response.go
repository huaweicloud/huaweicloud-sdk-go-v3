package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceTasksResponse Response Object
type ListInstanceTasksResponse struct {
	Page *Page `json:"page,omitempty"`

	// 实例任务列表
	Tasks          *[]InstanceTask `json:"tasks,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListInstanceTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTasksResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceTasksResponse", string(data)}, " ")
}
