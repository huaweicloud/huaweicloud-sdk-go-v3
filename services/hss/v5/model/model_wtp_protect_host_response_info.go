package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WtpProtectHostResponseInfo 网页防篡改防护信息
type WtpProtectHostResponseInfo struct {

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-256位，支持IPv4或IPv6格式（IPv4长度7-15位，IPv6长度15-39位）
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 策略组ID **取值范围**： 字符长度36-64位
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 服务器组名称 **取值范围**: 字符长度0-256位
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 资产重要性。 **取值范围**： - important ：重要资产。 - common ：一般资产。 - test ：测试资产。
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**： 操作系统位数 **取值范围**： 字符长度1-64位
	OsBit *string `json:"os_bit,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux - Windows：Windows
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**: 网页防篡改防护状态 **取值范围**: - opening : 开启中。 - opened : 防护中。 - closed : 未防护。 - open_failed : 防护失败。 - partial_protection : 部分防护。 - protection_interruption : 防护中断。 - protection_pause : 防护暂停。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**： 计费模式 **取值范围**： - packet_cycle ：包年/包月。 - on_demand ：按需计费。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释**： 资源ID **取值范围**： 字符长度0-128位
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**: 动态网页防篡改防护开启状态 **取值范围**: - opened ：已开启动态网页防篡改防护。 - closed ：未开启动态网页防篡改防护。
	RaspProtectStatus *string `json:"rasp_protect_status,omitempty"`

	// **参数解释**: 近7天静态网页防篡改攻击次数 **取值范围**: 最小值0，最大值2000000000
	AntiTamperingTimes *int64 `json:"anti_tampering_times,omitempty"`

	// **参数解释**: 近7天动态网页防篡改攻击次数 **取值范围**: 最小值0，最大值2000000000
	DetectTamperingTimes *int64 `json:"detect_tampering_times,omitempty"`

	// **参数解释**： 操作系统名称 **取值范围**： 字符长度0-128位
	OsName *string `json:"os_name,omitempty"`

	// **参数解释**： 操作系统版本 **取值范围**： 字符长度0-256位
	OsVersion *string `json:"os_version,omitempty"`

	// **参数解释**: 服务器状态 **取值范围**: - ACTIVE ：运行中。 - SHUTOFF ：关机。 - BUILDING ：创建中。 - ERROR ：故障。
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**: Agent状态 **取值范围**: - not_installed：未安装。 - online：在线。 - offline：离线。
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: 防护目录数 **取值范围**: 最小值0，最大值50
	ProtectDirNum *int32 `json:"protect_dir_num,omitempty"`

	// **参数解释**: 防护状态是防护失败的防护目录列表，仅当存在防护状态是防护失败的防护目录时返回。 **取值范围**: 最少0条，最多50条
	AbnormalDirList *[]string `json:"abnormal_dir_list,omitempty"`

	// **参数解释**: 防护失败原因，仅当存在防护状态是防护失败的防护目录时返回。 **取值范围**: - 1 ：某一个防护目录的路径不存在。 - 2 ：多个防护目录的路径不存在。 - 7 ：某一个防护目录未防护。 - 8 ：多个防护目录未防护。 - 10 ：某一个防护目录是网络文件系统。 - 11 ：多个防护目录是网络文件系统。
	AbnormalReason *int32 `json:"abnormal_reason,omitempty"`

	// **参数解释**： 远端备份服务器ID，仅当Linux服务器开启远端备份功能时返回。 **取值范围**： 字符长度0-64位
	BackupHostId *string `json:"backup_host_id,omitempty"`

	// **参数解释**： 防护中断原因，仅当防护状态是防护中断时返回。 **取值范围**： 字符长度0-256位
	InterruptReason *string `json:"interrupt_reason,omitempty"`
}

func (o WtpProtectHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WtpProtectHostResponseInfo struct{}"
	}

	return strings.Join([]string{"WtpProtectHostResponseInfo", string(data)}, " ")
}
