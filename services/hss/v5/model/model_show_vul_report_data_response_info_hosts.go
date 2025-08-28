package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVulReportDataResponseInfoHosts struct {

	// **参数解释**： 主机名称 **取值范围**： 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器公网ip **取值范围**： 字符长度1-64位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 服务器私网ip **取值范围**： 字符长度0-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 最近扫描时间 **取值范围**： 最小值0，最大值9223372036854775807
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**： 漏洞风险等级 **取值范围**： - Critical：紧急。 - High：高危。 - Medium：中危。 - Low：低危。
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**： 漏洞数 **取值范围**： 最小值0，最大值10000000
	VulTotalNum *int32 `json:"vul_total_num,omitempty"`

	// **参数解释**： 主机ID **取值范围**： 字符长度1-128位
	HostId *string `json:"host_id,omitempty"`
}

func (o ShowVulReportDataResponseInfoHosts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulReportDataResponseInfoHosts struct{}"
	}

	return strings.Join([]string{"ShowVulReportDataResponseInfoHosts", string(data)}, " ")
}
