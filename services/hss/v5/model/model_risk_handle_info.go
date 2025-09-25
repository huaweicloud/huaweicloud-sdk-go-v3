package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RiskHandleInfo 风险处置小结
type RiskHandleInfo struct {
	FreeReportInfo *RiskHandleInfoFreeReportInfo `json:"free_report_info,omitempty"`

	VulInfo *RiskHandleInfoVulInfo `json:"vul_info,omitempty"`

	BaseLineInfo *RiskHandleInfoBaseLineInfo `json:"base_line_info,omitempty"`

	AlarmInfo *RiskHandleInfoAlarmInfo `json:"alarm_info,omitempty"`
}

func (o RiskHandleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskHandleInfo struct{}"
	}

	return strings.Join([]string{"RiskHandleInfo", string(data)}, " ")
}
