package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowScanJobResultsResponse struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 任务名
	JobName *string `json:"job_name,omitempty" xml:"job_name"`

	// 查询资产类型
	Type *string `json:"type,omitempty" xml:"type"`

	DbScanResult *DbScanResult `json:"db_scan_result,omitempty" xml:"db_scan_result"`

	ObsScanResult *ObsScanResult `json:"obs_scan_result,omitempty" xml:"obs_scan_result"`

	EsScanResult   *EsScanResult `json:"es_scan_result,omitempty" xml:"es_scan_result"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowScanJobResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScanJobResultsResponse struct{}"
	}

	return strings.Join([]string{"ShowScanJobResultsResponse", string(data)}, " ")
}
