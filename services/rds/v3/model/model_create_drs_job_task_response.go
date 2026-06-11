package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrsJobTaskResponse Response Object
type CreateDrsJobTaskResponse struct {

	// 参数解释： 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDrsJobTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrsJobTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateDrsJobTaskResponse", string(data)}, " ")
}
