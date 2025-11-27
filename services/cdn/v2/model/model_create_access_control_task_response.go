package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccessControlTaskResponse Response Object
type CreateAccessControlTaskResponse struct {

	// 任务id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAccessControlTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessControlTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAccessControlTaskResponse", string(data)}, " ")
}
