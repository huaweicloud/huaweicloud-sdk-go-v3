package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTaskResponse Response Object
type DeleteTaskResponse struct {

	// 作业id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteTaskResponse", string(data)}, " ")
}
