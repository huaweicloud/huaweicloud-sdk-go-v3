package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowScanJobsResponse struct {
	// 本次返回的扫描任务列表

	Tasks *[]ScanJob `json:"tasks,omitempty"`
	// 任务总数

	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowScanJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScanJobsResponse struct{}"
	}

	return strings.Join([]string{"ShowScanJobsResponse", string(data)}, " ")
}
