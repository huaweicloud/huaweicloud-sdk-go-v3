package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisSummaryResponse Response Object
type ShowDiagnosisSummaryResponse struct {

	// error_code
	ErrorCode *string `json:"error_code,omitempty"`

	// error_msg
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 工单状态
	Status *string `json:"status,omitempty"`

	// region
	Region *string `json:"region,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 被诊断实例的概览信息列表
	InstanceSummary *[]DiagnosisSummaryItem `json:"instance_summary,omitempty"`
	HttpStatusCode  int                     `json:"-"`
}

func (o ShowDiagnosisSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisSummaryResponse", string(data)}, " ")
}
