package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFailedTaskResponse struct {

	// 失败任务总数。
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 失败的任务列表
	Tasks          *[]FailedTasks `json:"tasks,omitempty" xml:"tasks"`
	HttpStatusCode int            `json:"-"`
}

func (o ListFailedTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFailedTaskResponse struct{}"
	}

	return strings.Join([]string{"ListFailedTaskResponse", string(data)}, " ")
}
