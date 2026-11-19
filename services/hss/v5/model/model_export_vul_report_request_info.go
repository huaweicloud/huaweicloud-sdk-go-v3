package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportVulReportRequestInfo struct {

	// **参数解释**： 资产重要性 **约束限制**： 不涉及 **取值范围**： - important：重要资产 - common：一般资产 - test：测试资产  **默认取值**： 不涉及
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**： 服务器组名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位 **默认取值**： 不涉及
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 主机名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位 **默认取值**： 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 主机IP **约束限制**： 不涉及 **取值范围**： 字符长度1-64位 **默认取值**： 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**： 漏洞风险等级 **约束限制**： 不涉及 **取值范围**： - Security：安全 - High：高危 - Medium：中危 - Low：低危  **默认取值**： 不涉及
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**： 漏洞处理状态 **约束限制**： 必填 **取值范围**： - handled：已处理 - unhandled：未处理  **默认取值**： 不涉及
	HandleStatus string `json:"handle_status"`

	// **参数解释**： 漏洞状态 **约束限制**： 不涉及 **取值范围**： 字符长度1-32位 **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**： 主机ID列表 **约束限制**： 不涉及 **取值范围**： 0-1000个主机ID **默认取值**： 不涉及
	HostIds *[]string `json:"host_ids,omitempty"`

	// **参数解释**： 集群ID **约束限制**： 不涉及 **取值范围**： 字符长度1-64位 **默认取值**： 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o ExportVulReportRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportVulReportRequestInfo struct{}"
	}

	return strings.Join([]string{"ExportVulReportRequestInfo", string(data)}, " ")
}
