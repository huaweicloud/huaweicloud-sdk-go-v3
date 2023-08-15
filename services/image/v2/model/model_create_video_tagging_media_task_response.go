package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVideoTaggingMediaTaskResponse Response Object
type CreateVideoTaggingMediaTaskResponse struct {

	// 任务唯一标识
	TaskId *string `json:"task_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVideoTaggingMediaTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoTaggingMediaTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateVideoTaggingMediaTaskResponse", string(data)}, " ")
}
