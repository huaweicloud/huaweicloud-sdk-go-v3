package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNextflowTaskResponse Response Object
type ListNextflowTaskResponse struct {

	// 子任务实例
	Tasks *[]NextflowTaskListDto `json:"tasks,omitempty"`

	// 子任务的总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNextflowTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNextflowTaskResponse struct{}"
	}

	return strings.Join([]string{"ListNextflowTaskResponse", string(data)}, " ")
}
