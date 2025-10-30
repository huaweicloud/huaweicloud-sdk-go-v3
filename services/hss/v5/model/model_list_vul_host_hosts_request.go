package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulHostHostsRequest Request Object
type ListVulHostHostsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 资产重要性 **约束限制**: 不涉及 **取值范围**: - important ：重要资产 - common    ：一般资产 - test      ：测试资产  **默认取值**: 不涉及
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 服务器组 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 主机名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 主机公有ip **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 主机私有ip **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 主机风险等级 **约束限制**: 不涉及 **取值范围**: - High：高危 - Medium：中危 - Low：低危 - Security：安全  **默认取值**: 不涉及
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**: 主机的处置状态 **约束限制**: 不涉及 **取值范围**: - unhandled：待处理 - handled：已处理  **默认取值**: 不涉及
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**: 该漏洞状态包含的主机 **约束限制**: 不涉及 **取值范围**: - vul_status_unfix：未处理 - vul_status_ignored：忽略 - vul_status_verified：验证中 - vul_status_fixing：修复中 - vul_status_fixed：修复完成 - vul_status_reboot：修复成功待重启 - vul_status_failed：修复失败 - vul_status_fix_after_reboot：重启后再次修复  **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**: 集群id **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o ListVulHostHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHostHostsRequest struct{}"
	}

	return strings.Join([]string{"ListVulHostHostsRequest", string(data)}, " ")
}
