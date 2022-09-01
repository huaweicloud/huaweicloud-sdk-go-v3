package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBatchTaskResponse struct {
	Batchtask *Task `json:"batchtask,omitempty" xml:"batchtask"`

	// 子任务详情列表。
	TaskDetails *[]TaskDetail `json:"task_details,omitempty" xml:"task_details"`

	Page           *Page `json:"page,omitempty" xml:"page"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowBatchTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowBatchTaskResponse", string(data)}, " ")
}
