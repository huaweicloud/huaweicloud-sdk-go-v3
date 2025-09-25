package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportWebTamperHostRequestInfo 导出数据的表头字段列表
type ExportWebTamperHostRequestInfo struct {

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器ID列表 **取值范围**: 不涉及
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-256位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: Agent ID。（已废弃，请使用host_id） **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux。 - Windows：Windows。
	OsType *string `json:"os_type,omitempty"`

	// 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 服务器组名称 **取值范围**: 字符长度0-256位
	GroupName *string `json:"group_name,omitempty"`

	// 防护状态。（已废弃，请使用wtp_status）
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**: Agent状态。 **约束限制**: 不涉及 **取值范围**: - not_installed：未安装。 - online：在线。 - offline：离线。  **默认取值**: 不涉及。
	AgentStatus *string `json:"agent_status,omitempty"`

	// 动态网页防篡改防护状态   - opened : 防护中   - closed : 未防护
	RaspProtectStatus *string `json:"rasp_protect_status,omitempty"`

	// **参数解释**: 网页防篡改防护状态 **取值范围**: - opening : 开启中。 - opened : 防护中。 - open_failed : 防护失败。 - partial_protection : 部分防护。 - protection_interruption : 防护中断。
	WtpStatus *string `json:"wtp_status,omitempty"`

	// 偏移量。（已废弃，请使用export_size）
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数。（已废弃，请使用export_size）
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 导出数据表头信息详情 **约束限制**: 表头信息应为如下格式[[字段1,表头1显示名称],[字段2,表头2显示名称],[字段3,表头3显示名称]] **取值范围**: 可从如下取值中选取部分或全部组成表头信息 - host_name：服务器名称。 - host_id：服务器ID。 - public_ip：弹性公网IP地址。 - private_ip：私有IP地址。 - charging_mode：计费模式。 - agent_status：Agent状态。 - protect_status：防护状态。 - protect_dir_num：防护目录数。 - rasp_protect_status：动态防篡改状态。 - anti_tampering_times：静态防篡改攻击（近七天）。 - detect_tampering_times：动态防篡改攻击（近七天）。 - protect_dir：防护目录。 - exclude_child_dir：排除子目录。 - exclude_file_list：排除文件路径列表。 - local_backup_dir：本地备份路径。 - exclude_file_type：排除文件类型。 - dir_protect_status：防护目录状态。 - error：错误信息。  **默认取值**: 不涉及
	ExportHeaders [][]string `json:"export_headers"`

	// **参数解释**: 导出数据条数上限 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值200000 **默认取值**: 不涉及
	ExportSize int32 `json:"export_size"`
}

func (o ExportWebTamperHostRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportWebTamperHostRequestInfo struct{}"
	}

	return strings.Join([]string{"ExportWebTamperHostRequestInfo", string(data)}, " ")
}
