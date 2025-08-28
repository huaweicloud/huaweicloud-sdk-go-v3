package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulReportRequestBody struct {

	// **参数解释**: 资产重要性 **约束限制**: 不涉及 **取值范围**: - important：重要资产。 - common：一般资产。 - test：测试资产。  **默认取值**: 不涉及
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**： 服务器组名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-64位 **默认取值**： 不涉及
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 服务器ID列表 **约束限制**： 不涉及 **取值范围**： 0-100个items **默认取值**： 不涉及
	HostIds *[]string `json:"host_ids,omitempty"`

	// **参数解释**： 主机名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-64位 **默认取值**： 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 主机ip **约束限制**： 不涉及 **取值范围**： 字符长度0-128位 **默认取值**： 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 漏洞风险等级 **约束限制**: 不涉及 **取值范围**: - High：高。 - Medium：中等。 - Low：低。 - Security：安全。  **默认取值**: 不涉及
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**: 处置状态 **约束限制**: 不涉及 **取值范围**: - unhandled：未处理。 - handled：已处理。  **默认取值**: 不涉及
	HandleStatus string `json:"handle_status"`

	// **参数解释**: 主机的漏洞状态 **约束限制**: 不涉及 **取值范围**: - vul_status_unfix：未处理。 - vul_status_ignored：已忽略。 - vul_status_verified：验证中。 - vul_status_fixing：修复中。 - vul_status_fixed：已修复。 - vul_status_reboot：修复待重启。 - vul_status_failed：修复失败。 - vul_status_fix_after_reboot：请重启主机再次修复。  **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`
}

func (o VulReportRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulReportRequestBody struct{}"
	}

	return strings.Join([]string{"VulReportRequestBody", string(data)}, " ")
}
