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

	// **参数解释**: 策略支持的操作系统 **约束限制**: 不涉及 **取值范围**: 包含如下：   - Windows : Windows系统   - Linux : Linux系统 **默认取值**: 不涉及
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**: 应用防护状态 **约束限制**: 不涉及 **取值范围**: 包含如下7种。 - app_install_processing：防护开启中。 - app_protected：防护成功。 - app_install_failed：防护失败（安装失败）。 - app_not_configure：未防护。 - app_partially_protected：部分防护。 - app_all_failed：防护失败。 - app_uninstall_processing：卸载中。 **默认取值**: 不涉及
	RaspStatus *string `json:"rasp_status,omitempty"`

	// 防护策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 是否为友好用户
	IsFriendlyUser *bool `json:"is_friendly_user,omitempty"`

	// agent是否支持动态加载
	AgentSupportAutoAttach *bool `json:"agent_support_auto_attach,omitempty"`

	// **参数解释**: Agent状态 **约束限制**: 不涉及 **取值范围**: 包含如下6种。 - installed ：已安装。 - not_installed ：未安装。 - online ：在线。 - offline ：离线。 - install_failed ：安装失败。 - installing ：安装中。 **默认取值**: 不涉及
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
