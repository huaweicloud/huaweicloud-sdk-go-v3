package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTaskListByProjectIdResponse struct {

	// 任务信息
	Tasks *[]SimpleTaskInfoV2 `json:"tasks,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTaskListByProjectIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskListByProjectIdResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskListByProjectIdResponse", string(data)}, " ")
}
