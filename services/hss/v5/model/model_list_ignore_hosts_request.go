package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIgnoreHostsRequest Request Object
type ListIgnoreHostsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 主机开通的版本 **约束限制**: 不涉及 **取值范围**： 包含如下7种输入。 - hss.version.null ：无。 - hss.version.basic ：基础版。 - hss.version.advanced ：专业版。 - hss.version.enterprise ：企业版。 - hss.version.premium ：旗舰版。 - hss.version.wtp ：网页防篡改版。 - hss.version.container.enterprise：容器版。 - hss.version.container.wtp：容器网页防篡改增值项。  **默认取值**: 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**： Agent状态 **约束限制**: 不涉及 **取值范围**： - installed：已安装 - not_installed：未安装 - online：在线 - offline：离线 - install_failed：安装失败 - installing：安装中  **默认取值**: 不涉及
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: 检测结果 **约束限制**: 不涉及 **取值范围**: 包含如下4种： - undetected：未检测 - clean：无风险 - risk：有风险 - scanning：检测中  **默认取值**: 不涉及
	DetectResult *string `json:"detect_result,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器的唯一标识ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 主机状态 **约束限制**: 不涉及 **取值范围**: 包含如下4种： - ACTIVE：正在运行 - SHUTOFF：关机 - BUILDING：创建中 - ERROR：故障  **默认取值**: 不涉及
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**: 操作系统类型。 **约束限制**: 不涉及 **取值范围**: - Linux：Linux操作系统。 - Windows：Windows操作系统。  **默认取值**: 无
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**: 服务器私有IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器弹性IP地址 **约束限制**: 不涉及 **取值范围**: IPv4格式（长度7-15位）、IPv6格式（长度15-39位） **默认取值**: 无
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 公网或私网IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	IpAddr *string `json:"ip_addr,omitempty"`

	// **参数解释**: 防护状态 **约束限制**: 不涉及 **取值范围**: 包含如下3种： - closed：关闭 - opened：开启 - protection_exception：防护异常  **默认取值**: 不涉及
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**: 服务器组的唯一标识ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 服务器组名称。 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 无
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 策略组ID **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// **参数解释**: 策略组名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// **参数解释**： 收费模式 **约束限制**: 不涉及 **取值范围**: - packet_cycle ：包年/包月。 - on_demand ：按需。  **默认取值**: 不涉及
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释**: 资产重要性 **约束限制**: 不涉及 **取值范围**： - important：重要资产 - common：一般资产 - test：测试资产  **默认取值**: 不涉及
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 服务器标签，格式为\"key:value\" **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	Label *string `json:"label,omitempty"`

	// **参数解释**: 存在告警事件 **约束限制**: 不涉及 **取值范围**: true或者false **默认取值**: 不涉及
	HasIntrusion *bool `json:"has_intrusion,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListIgnoreHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIgnoreHostsRequest struct{}"
	}

	return strings.Join([]string{"ListIgnoreHostsRequest", string(data)}, " ")
}
