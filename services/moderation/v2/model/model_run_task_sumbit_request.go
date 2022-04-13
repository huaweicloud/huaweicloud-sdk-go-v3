package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunTaskSumbitRequest struct {
	Body *TaskSumbitReq `json:"body,omitempty"`
}

func (o RunTaskSumbitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTaskSumbitRequest struct{}"
	}

	return strings.Join([]string{"RunTaskSumbitRequest", string(data)}, " ")
}
