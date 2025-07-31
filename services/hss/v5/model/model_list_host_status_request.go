package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostStatusRequest Request Object
type ListHostStatusRequest struct {

	// Region ID
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 主机开通的版本 **约束限制**: 不涉及 **取值范围**： 包含如下7种输入。 - hss.version.null ：无。 - hss.version.basic ：基础版。 - hss.version.advanced ：专业版。 - hss.version.enterprise ：企业版。 - hss.version.premium ：旗舰版。 - hss.version.wtp ：网页防篡改版。 - hss.version.container.enterprise：容器版。 **默认取值**: 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**: Agent的状态 **约束限制**: 不涉及 **取值范围**: Agent的状态分为两类： - installed：已安装。已安装状态包含以下四种情况：   - online：在线。表示Agent已经成功安装并且与HSS云端防护中心保持连接。   - offline：离线。表示虽然Agent已经安装，但当前与HSS云端防护中心的连接中断。   - install_failed：安装失败。表示在尝试安装过程中遇到错误或问题，导致安装未能完成。   - installing：安装中。表示当前正在进行Agent安装。 - not_installed ：未安装。表示服务器中尚未安装Agent。 如果您要筛选除在线以外所有状态的Agent，可设置not_online（仅作为查询条件） **默认取值**: 不涉及
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: 检测结果 **约束限制**: 不涉及 **取值范围**: 包含如下4种。   - undetected ：未检测。   - clean ：无风险。   - risk ：有风险。   - scanning ：检测中。 **默认取值**: 不涉及
	DetectResult *string `json:"detect_result,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 主机状态 **约束限制**: 不涉及 **取值范围**: 包含如下4种。 - ACTIVE ：正在运行。 - SHUTOFF ：关机。 - BUILDING ：创建中。 - ERROR ：故障。 **默认取值**: 不涉及
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**: 操作系统类型 **约束限制**: 不涉及 **取值范围**: 包含如下2种。 - Linux ：Linux。 - Windows ：Windows。 **默认取值**: 不涉及
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**: 服务器私有IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// 服务器公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 公网或私网IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	IpAddr *string `json:"ip_addr,omitempty"`

	// **参数解释**: 防护状态 **约束限制**: 不涉及 **取值范围**: 包含如下3种： - closed ：关闭。 - opened ：开启。 - protection_exception ：防护异常。 **默认取值**: 不涉及
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 服务器组ID
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 服务器组名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: VPC的ID **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**: 存在告警事件 **约束限制**: 不涉及 **取值范围**: true或者false **默认取值**: 不涉及
	HasIntrusion *bool `json:"has_intrusion,omitempty"`

	// **参数解释**: 存在漏洞风险 **约束限制**: 不涉及 **取值范围**: true或者false **默认取值**: 不涉及
	HasVul *bool `json:"has_vul,omitempty"`

	// **参数解释**: 存在基线风险 **约束限制**: 不涉及 **取值范围**: true或者false **默认取值**: 不涉及
	HasBaseline *bool `json:"has_baseline,omitempty"`

	// **参数解释**: 排序字段 **约束限制**: 不涉及 **取值范围**: 只支持risk_num **默认取值**: 不涉及
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**: 排序的顺序 **约束限制**: 不涉及 **取值范围**: - asc: 正序 - desc: 倒序 **默认取值**: 不涉及
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**: 策略组ID **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// **参数解释**: 策略组名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// **参数解释**： 收费模式 **约束限制**: 不涉及 **取值范围**: - packet_cycle ：包年/包月。 - on_demand ：按需。 **默认取值**: 不涉及
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释** : 是否强制从ECS同步主机 **约束限制** : 不涉及 **取值范围** : true或者false **默认取值** : 不涉及
	Refresh *bool `json:"refresh,omitempty"`

	// **参数解释** : 是否获取主机常用登录地信息 **约束限制** : 不涉及 **取值范围** : true或者false **默认取值** : 不涉及
	GetCommonLoginLocations *bool `json:"get_common_login_locations,omitempty"`

	// **参数解释** : 是否返回比当前版本高的所有版本 **约束限制** : 不涉及 **取值范围** : true或者false **默认取值** : 不涉及
	AboveVersion *bool `json:"above_version,omitempty"`

	// **参数解释** : 是否华为云主机 **约束限制** : 不涉及 **取值范围** : true或者false **默认取值** : 不涉及
	OutsideHost *bool `json:"outside_host,omitempty"`

	// **参数解释** : 资产重要性 **约束限制** : 不涉及 **取值范围** : 包含如下4种 - important ：重要资产 - common ：一般资产 - test ：测试资产 **默认取值** : 不涉及
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释** : 资产标签 **约束限制** : 不涉及 **取值范围** : 字符长度1-64位 **默认取值** : 不涉及
	Label *string `json:"label,omitempty"`

	// **参数解释** : 资产服务器组 **约束限制** : 不涉及 **取值范围** : 字符长度1-64位 **默认取值** : 不涉及
	ServerGroup *string `json:"server_group,omitempty"`

	// **参数解释** : agent是否可升级 **约束限制** : 不涉及 **取值范围** : true或者false **默认取值** : 不涉及
	AgentUpgradable *bool `json:"agent_upgradable,omitempty"`

	// **参数解释** : 是否安装模式场景 **约束限制** : 不涉及 **取值范围** : true或者false **默认取值** : 不涉及
	InstallMode *bool `json:"install_mode,omitempty"`

	// **参数解释** : 是否绑定DEW密钥 **约束限制** : 不涉及 **取值范围** : true或者false **默认取值** : 不涉及
	BindingKey *bool `json:"binding_key,omitempty"`

	// **参数解释** : 是否防护中断 **约束限制** : 不涉及 **取值范围** : true或者false **默认取值** : 不涉及
	ProtectInterrupt *bool `json:"protect_interrupt,omitempty"`

	// **参数解释** : 是否集群内节点 **约束限制** : 不涉及 **取值范围** : true或者false **默认取值** : 不涉及
	Incluster *bool `json:"incluster,omitempty"`

	// **参数解释** : 是否防护降级 **约束限制** : 不涉及 **取值范围** : true或者false **默认取值** : 不涉及
	ProtectDegradation *bool `json:"protect_degradation,omitempty"`

	// **参数解释**: 集群ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListHostStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostStatusRequest struct{}"
	}

	return strings.Join([]string{"ListHostStatusRequest", string(data)}, " ")
}
