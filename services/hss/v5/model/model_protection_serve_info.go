package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectionServeInfo struct {

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// Agent ID
	AgentId *string `json:"agent_id,omitempty"`

	// Agent版本
	AgentVersion *string `json:"agent_version,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 弹性ip地址
	PublicIp *string `json:"public_ip,omitempty"`

	// 私有ip
	PrivateIp *string `json:"private_ip,omitempty"`

	// 操作系统类型
	OsType *string `json:"os_type,omitempty"`

	// 应用防护状态 |- 应用防护状态，包含如下6种。 - 0 ：防护开启中。 - 2 ：防护成功。 - 4 ：防护失败（安装失败）。 - 6 ：未防护。 - 8 ：部分防护。 - 9 ：防护失败。
	RaspStatus *string `json:"rasp_status,omitempty"`

	// 防护策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 是否为友好用户
	IsFriendlyUser *bool `json:"is_friendly_user,omitempty"`

	// agent是否支持动态加载
	AgentSupportAutoAttach *bool `json:"agent_support_auto_attach,omitempty"`

	// Agent状态
	AgentStatus *string `json:"agent_status,omitempty"`

	// 动态加载是否开启
	AutoAttach *bool `json:"auto_attach,omitempty"`

	// 防护状态 |- agent防护状态，包含如下2种。 - 0 ：关闭。 - 1 ：开启。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 服务器组ID
	GroupId *string `json:"group_id,omitempty"`

	// 服务器组名称
	GroupName *string `json:"group_name,omitempty"`

	// 防护事件数
	ProtectEventNum *int64 `json:"protect_event_num,omitempty"`

	// RASP端口号
	RaspPort *int32 `json:"rasp_port,omitempty"`
}

func (o ProtectionServeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionServeInfo struct{}"
	}

	return strings.Join([]string{"ProtectionServeInfo", string(data)}, " ")
}
