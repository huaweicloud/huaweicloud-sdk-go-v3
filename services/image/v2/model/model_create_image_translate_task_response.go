package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateImageTranslateTaskResponse struct {

	// 任务唯一标识
	TaskId *string `json:"task_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateImageTranslateTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageTranslateTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateImageTranslateTaskResponse", string(data)}, " ")
}
