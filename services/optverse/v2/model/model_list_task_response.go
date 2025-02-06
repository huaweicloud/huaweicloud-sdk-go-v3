package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskResponse Response Object
type ListTaskResponse struct {

	// 任务数量
	Count *int32 `json:"count,omitempty"`

	// 任务列表项视图
	Tasks          *[]TaskListItemVo `json:"tasks,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskResponse struct{}"
	}

	return strings.Join([]string{"ListTaskResponse", string(data)}, " ")
}
