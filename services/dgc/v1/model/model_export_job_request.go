package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportJobRequest struct {
	// 作业名称.

	JobName string `json:"job_name"`
}

func (o ExportJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportJobRequest struct{}"
	}

	return strings.Join([]string{"ExportJobRequest", string(data)}, " ")
}
