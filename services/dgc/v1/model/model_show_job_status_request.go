package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJobStatusRequest struct {
	// 作业名称.

	JobName string `json:"job_name"`
}

func (o ShowJobStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowJobStatusRequest", string(data)}, " ")
}
