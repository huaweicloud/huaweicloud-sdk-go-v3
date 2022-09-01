package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowScanJobsResponse struct {

	// 本次返回的扫描任务列表
	Tasks *[]ScanJob `json:"tasks,omitempty" xml:"tasks"`

	// 任务总数
	Total          *int64 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowScanJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScanJobsResponse struct{}"
	}

	return strings.Join([]string{"ShowScanJobsResponse", string(data)}, " ")
}
