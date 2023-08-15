package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteActionResponse Response Object
type ExecuteActionResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteActionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteActionResponse", string(data)}, " ")
}
