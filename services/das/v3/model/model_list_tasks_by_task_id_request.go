package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksByTaskIdRequest Request Object
type ListTasksByTaskIdRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`

	// 页码
	CurPage *int32 `json:"cur_page,omitempty"`

	// 每页记录数
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ListTasksByTaskIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksByTaskIdRequest struct{}"
	}

	return strings.Join([]string{"ListTasksByTaskIdRequest", string(data)}, " ")
}
