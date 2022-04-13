package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportDatasetRequest struct {
	// 作业ID。

	JobId string `json:"job_id"`
	// 作业运行ID。

	RunId string `json:"run_id"`
}

func (o ExportDatasetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDatasetRequest struct{}"
	}

	return strings.Join([]string{"ExportDatasetRequest", string(data)}, " ")
}
