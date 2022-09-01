package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateTaskResponse struct {

	// 任务ID。
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 任务名称。
	TaskName       *string `json:"task_name,omitempty" xml:"task_name"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateTaskResponse", string(data)}, " ")
}
