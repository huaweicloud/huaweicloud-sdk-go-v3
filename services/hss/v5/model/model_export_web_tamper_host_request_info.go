package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportWebTamperHostRequestInfo 导出数据的表头字段列表
type ExportWebTamperHostRequestInfo struct {

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// 主机ID数组
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-256位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: Agent ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux。 - Windows：Windows。
	OsType *string `json:"os_type,omitempty"`

	// 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 服务器组名称 **取值范围**: 字符长度0-256位
	GroupName *string `json:"group_name,omitempty"`

	// 防护状态，包含如下2种。   - closed ：关闭。   - opened ：开启。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// Agent状态，包含如下6种。   - installed ：已安装。   - not_installed ：未安装。   - online ：在线。   - offline ：离线。   - install_failed ：安装失败。   - installing ：安装中。   - not_online ：不在线的（除了在线以外的所有状态，仅作为查询条件）。
	AgentStatus *string `json:"agent_status,omitempty"`

	// 动态网页防篡改防护状态   - opened : 防护中   - closed : 未防护
	RaspProtectStatus *string `json:"rasp_protect_status,omitempty"`

	// 防护状态   - closed : 未防护   - opened : 防护中   - open_failed : 防护失败   - opening : 正在开启   - partial_protection : 部分防护
	WtpStatus *string `json:"wtp_status,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 导出数据条数
	ExportSize *int32 `json:"export_size,omitempty"`

	// 导出表头集合
	ExportHeaders *[][]string `json:"export_headers,omitempty"`
}

func (o ExportWebTamperHostRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportWebTamperHostRequestInfo struct{}"
	}

	return strings.Join([]string{"ExportWebTamperHostRequestInfo", string(data)}, " ")
}
