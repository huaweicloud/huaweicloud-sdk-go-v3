package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartJobRequest struct {
	// 作业名称.

	JobName string `json:"job_name"`

	Body *StartJobReq `json:"body,omitempty"`
}

func (o StartJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartJobRequest struct{}"
	}

	return strings.Join([]string{"StartJobRequest", string(data)}, " ")
}
