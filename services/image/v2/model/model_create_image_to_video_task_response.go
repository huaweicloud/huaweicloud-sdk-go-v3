package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateImageToVideoTaskResponse struct {

	// 任务唯一标识
	TaskId *string `json:"task_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateImageToVideoTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageToVideoTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateImageToVideoTaskResponse", string(data)}, " ")
}
