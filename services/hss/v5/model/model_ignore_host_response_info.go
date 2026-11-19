package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IgnoreHostResponseInfo 已忽略主机信息
type IgnoreHostResponseInfo struct {

	// **参数解释**: 主机上安装的杀毒Agent的唯一标识ID，用于关联主机与杀毒服务 **取值范围**: 字符长度1-64位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 服务器的唯一标识ID **取值范围**: 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器私有IP **取值范围**: 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器弹性IP地址 **取值范围**: IPv4格式（长度7-15位）、IPv6格式（长度15-39位）
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 操作系统名称 **取值范围**: 字符长度1-64位
	OsName *string `json:"os_name,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux - Windows：Windows
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 服务器状态 **取值范围**： - ACTIVE：运行中 - SHUTOFF：关机 - BUILDING：创建中 - ERROR：故障
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**： Agent状态 **取值范围**： - installed：已安装 - not_installed：未安装 - online：在线 - offline：离线 - install_failed：安装失败 - installing：安装中
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**： agent版本 **取值范围**： 字符长度1-32位
	AgentVersion *string `json:"agent_version,omitempty"`

	// **参数解释**: 服务器组ID **取值范围**: 字符范围1-64位
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 服务器组名称 **取值范围**: 字符范围1-128位
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 资产重要性 **约束限制**: 不涉及 **取值范围**： - important：重要资产 - common：一般资产 - test：测试资产  **默认取值**: 不涉及
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释** 企业项目ID **取值范围** 字符长度0-64位
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** 企业项目名称 **取值范围** 字符长度0-256位
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// **参数解释** 操作系统位数 **取值范围** 字符长度0-64位
	OsBit *string `json:"os_bit,omitempty"`

	// **参数解释** 操作系统版本 **取值范围** 字符长度0-64位
	OsVersion *string `json:"os_version,omitempty"`

	// **参数解释** 防护状态 **取值范围** - closed：关闭 - opened：开启 - protection_exception：防护异常
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释** 检测结果 **取值范围** - undetected：未检测 - clean：无风险 - risk：有风险 - scanning：检测中
	DetectResult *string `json:"detect_result,omitempty"`

	// **参数解释** 计费模式 **取值范围** 字符长度0-32位
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释** 主机资源ID **取值范围** 字符长度0-128位
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释** 策略组ID **取值范围** 字符长度0-128位
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// **参数解释** 策略组名称 **取值范围** 字符长度0-256位
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// **参数解释** 主机防护版本到期时间，毫秒级时间戳 **取值范围** 0-9223372036854775807
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// **参数解释** 开通防护时间，毫秒级时间戳 **取值范围** 0-9223372036854775807
	OpenTime *int64 `json:"open_time,omitempty"`

	// **参数解释** Agent安装时间，毫秒级时间戳 **取值范围** 0-9223372036854775807
	AgentCreateTime *int64 `json:"agent_create_time,omitempty"`

	// **参数解释** Agent更新时间，毫秒级时间戳 **取值范围** 0-9223372036854775807
	AgentUpdateTime *int64 `json:"agent_update_time,omitempty"`

	// **参数解释** Agent是否可升级 **取值范围** true或者false
	Upgradable *bool `json:"upgradable,omitempty"`

	// **参数解释** 是否非本账号主机 **取值范围** true或者false
	OutsideHost *bool `json:"outside_host,omitempty"`

	// **参数解释** 容器类型 **取值范围** 字符长度0-32位
	ContainerType *string `json:"container_type,omitempty"`

	// **参数解释** 网页防篡改防护状态 **取值范围** - closed：关闭 - opened：开启
	WtpProtectStatus *string `json:"wtp_protect_status,omitempty"`
}

func (o IgnoreHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IgnoreHostResponseInfo struct{}"
	}

	return strings.Join([]string{"IgnoreHostResponseInfo", string(data)}, " ")
}
