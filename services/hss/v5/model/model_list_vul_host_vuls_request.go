package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulHostVulsRequest Request Object
type ListVulHostVulsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 漏洞名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: 漏洞cve编号 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释**: 漏洞标签 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	LabelList *string `json:"label_list,omitempty"`

	// **参数解释**: 漏洞状态 **约束限制**: 不涉及 **取值范围**: - vul_status_unfix：未处理 - vul_status_ignored：已忽略 - vul_status_verified：验证中 - vul_status_fixing：修复中 - vul_status_fixed：修复成功 - vul_status_reboot：修复成功待重启 - vul_status_failed：修复失败 - vul_status_fix_after_reboot：请重启主机再次修复  **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**: 主机的资产重要性 **约束限制**: 不涉及 **取值范围**: - important：重要资产 - common：一般资产 - test：测试资产  **默认取值**: 不涉及
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 主机的所属服务器组 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 主机名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 主机ip **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 查询的漏洞类型 **约束限制**: 不涉及 **取值范围**: - linux_vul：linux漏洞 - windows_vul：windows漏洞 - web_cms：Web-CMS漏洞 - app_vul：应用漏洞 - urgent_vul：应急漏洞  **默认取值**: 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**: 查询漏洞的修复紧急程度类型 **约束限制**: 不涉及 **取值范围**: - need_urgent_repair：需要紧急修复的漏洞 - unrepair：未完成修复的漏洞  **默认取值**: 不涉及
	RepairType *string `json:"repair_type,omitempty"`

	// **参数解释**: 漏洞风险等级 **约束限制**: 不涉及 **取值范围**: - Critical: 紧急 - High: 高 - Medium: 中 - Low: 低  **默认取值**: 不涉及
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**: 漏洞修复优先级 **约束限制**: 不涉及 **取值范围**: - Critical：紧急 - High：高 - Medium：中 - Low：低 **默认取值**: 不涉及
	RepairPriority *string `json:"repair_priority,omitempty"`
}

func (o ListVulHostVulsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHostVulsRequest struct{}"
	}

	return strings.Join([]string{"ListVulHostVulsRequest", string(data)}, " ")
}
