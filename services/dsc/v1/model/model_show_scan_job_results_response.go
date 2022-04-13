package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowScanJobResultsResponse struct {
	// 任务ID

	JobId *string `json:"job_id,omitempty"`
	// 任务名

	JobName *string `json:"job_name,omitempty"`
	// 查询资产类型

	Type *string `json:"type,omitempty"`

	DbScanResult *DbScanResult `json:"db_scan_result,omitempty"`

	ObsScanResult *ObsScanResult `json:"obs_scan_result,omitempty"`

	EsScanResult   *EsScanResult `json:"es_scan_result,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowScanJobResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScanJobResultsResponse struct{}"
	}

	return strings.Join([]string{"ShowScanJobResultsResponse", string(data)}, " ")
}
