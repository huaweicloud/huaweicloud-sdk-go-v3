package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulReportDataResponse Response Object
type ShowVulReportDataResponse struct {
	Sumary *ShowVulReportDataResponseInfoSumary `json:"sumary,omitempty"`

	// 主机列表
	Hosts *[]ShowVulReportDataResponseInfoHosts `json:"hosts,omitempty"`

	// 漏洞列表
	Vulnerabilities *[]ShowVulReportDataResponseInfoVulnerabilities `json:"vulnerabilities,omitempty"`

	// **参数解释**： 报告生成时间 **取值范围**： 最小值0，最大值9223372036854775807
	ReportCreateTime *int64 `json:"report_create_time,omitempty"`

	// **参数解释**： 漏洞总数量 **取值范围**： 最小值0，最大值10000000
	VulnerabilityTotalCount *int32 `json:"vulnerability_total_count,omitempty"`

	// **参数解释**： 主机总数量 **取值范围**： 最小值0，最大值20000
	HostTotalCount *int32 `json:"host_total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVulReportDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulReportDataResponse struct{}"
	}

	return strings.Join([]string{"ShowVulReportDataResponse", string(data)}, " ")
}
