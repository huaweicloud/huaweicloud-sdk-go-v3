package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksResponse Response Object
type ListTasksResponse struct {

	// 作业数量
	Count *int32 `json:"count,omitempty"`

	// 作业详情
	Tasks          *[]TaskInfo `json:"tasks,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksResponse struct{}"
	}

	return strings.Join([]string{"ListTasksResponse", string(data)}, " ")
}
