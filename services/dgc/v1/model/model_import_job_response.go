package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ImportJobResponse struct {
	// 任务id

	TaskId         *string `json:"taskId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportJobResponse struct{}"
	}

	return strings.Join([]string{"ImportJobResponse", string(data)}, " ")
}
