package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostStatusRequest Request Object
type ListHostStatusRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 主机开通的版本，包含如下7种输入。   - hss.version.null ：无。   - hss.version.basic ：基础版。   - hss.version.advanced ：专业版。   - hss.version.enterprise ：企业版。   - hss.version.premium ：旗舰版。   - hss.version.wtp ：网页防篡改版。   - hss.version.container.enterprise：容器版。
	Version *string `json:"version,omitempty"`

	// Agent的状态分为两类： - installed：已安装。已安装状态包含以下四种情况：    - online：在线。表示Agent已经成功安装并且与HSS云端防护中心保持连接。    - offline：离线。表示虽然Agent已经安装，但当前与HSS云端防护中心的连接中断。    - install_failed：安装失败。表示在尝试安装过程中遇到错误或问题，导致安装未能完成。    - installing：安装中。表示当前正在进行Agent安装。 - not_installed ：未安装。表示服务器中尚未安装Agent。 如果您要筛选除在线以外所有状态的Agent，可设置not_online（仅作为查询条件）
	AgentStatus *string `json:"agent_status,omitempty"`

	// 检测结果，包含如下4种。   - undetected ：未检测。   - clean ：无风险。   - risk ：有风险。   - scanning ：检测中。
	DetectResult *string `json:"detect_result,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 主机状态，包含如下4种。   - ACTIVE ：正在运行。   - SHUTOFF ：关机。   - BUILDING ：创建中。   - ERROR ：故障。
	HostStatus *string `json:"host_status,omitempty"`

	// 操作系统类型，包含如下2种。   - Linux ：Linux。   - Windows ：Windows。
	OsType *string `json:"os_type,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 服务器公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 公网或私网IP
	IpAddr *string `json:"ip_addr,omitempty"`

	// 防护状态，包含如下2种。   - closed ：关闭。   - opened ：开启。   - protection_exception ：防护异常。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 服务器组ID
	GroupId *string `json:"group_id,omitempty"`

	// 服务器组名称
	GroupName *string `json:"group_name,omitempty"`

	// vpc id
	VpcId *string `json:"vpc_id,omitempty"`

	// Region ID
	Region *string `json:"region,omitempty"`

	// 存在告警事件
	HasIntrusion *bool `json:"has_intrusion,omitempty"`

	// 存在漏洞风险
	HasVul *bool `json:"has_vul,omitempty"`

	// 存在基线风险
	HasBaseline *bool `json:"has_baseline,omitempty"`

	// 排序字段，只支持risk_num - risk_num：风险总量
	SortKey *string `json:"sort_key,omitempty"`

	// 排序的顺序 - asc: 正序 - desc: 倒序
	SortDir *string `json:"sort_dir,omitempty"`

	// 策略组ID
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// 策略组名称
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// 收费模式，包含如下2种。   - packet_cycle ：包年/包月。   - on_demand ：按需。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 是否强制从ECS同步主机
	Refresh *bool `json:"refresh,omitempty"`

	// 是否获取主机常用登录地信息
	GetCommonLoginLocations *bool `json:"get_common_login_locations,omitempty"`

	// 是否返回比当前版本高的所有版本
	AboveVersion *bool `json:"above_version,omitempty"`

	// 是否华为云主机
	OutsideHost *bool `json:"outside_host,omitempty"`

	// 资产重要性，包含如下4种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// 资产标签
	Label *string `json:"label,omitempty"`

	// 资产服务器组
	ServerGroup *string `json:"server_group,omitempty"`

	// agent是否可升级
	AgentUpgradable *bool `json:"agent_upgradable,omitempty"`

	// 是否安装模式场景
	InstallMode *bool `json:"install_mode,omitempty"`

	// 是否绑定DEW密钥
	BindingKey *bool `json:"binding_key,omitempty"`

	// 是否防护中断
	ProtectInterrupt *bool `json:"protect_interrupt,omitempty"`

	// 是否集群内节点
	Incluster *bool `json:"incluster,omitempty"`

	// 是否防护降级
	ProtectDegradation *bool `json:"protect_degradation,omitempty"`

	// 每页显示数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListHostStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostStatusRequest struct{}"
	}

	return strings.Join([]string{"ListHostStatusRequest", string(data)}, " ")
}
