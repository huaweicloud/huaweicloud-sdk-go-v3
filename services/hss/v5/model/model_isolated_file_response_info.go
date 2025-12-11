package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsolatedFileResponseInfo 已隔离文件详情
type IsolatedFileResponseInfo struct {

	// **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux - Windows：Windows
	OsType string `json:"os_type"`

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId string `json:"host_id"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName string `json:"host_name"`

	// **参数解释**： 文件哈希 **取值范围**： 字符长度1-256位
	FileHash string `json:"file_hash"`

	// **参数解释**： 文件路径 **取值范围**： 字符长度1-256位
	FilePath string `json:"file_path"`

	// **参数解释**： 文件的系统属性（如读写权限、隐藏属性、执行权限等） **取值范围**： 字符长度1-256位
	FileAttr string `json:"file_attr"`

	// 隔离状态，包含如下:   - isolated : 已隔离   - restored : 已恢复   - isolating : 已下发隔离任务   - restoring : 已下发恢复任务
	IsolationStatus string `json:"isolation_status"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp string `json:"private_ip"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-256位，支持IPv4或IPv6格式（IPv4长度7-15位，IPv6长度15-39位）
	PublicIp string `json:"public_ip"`

	// **参数解释**： 资产重要性。 **取值范围**： - important ：重要资产。 - common ：一般资产。 - test ：测试资产。
	AssetValue string `json:"asset_value"`

	// 更新时间，毫秒
	UpdateTime int64 `json:"update_time"`

	// agent版本
	AgentVersion string `json:"agent_version"`

	// 隔离来源，包含如下:   - event : 安全告警事件   - antivirus : 病毒查杀
	IsolateSource string `json:"isolate_source"`

	// **参数解释**： 事件名称 **取值范围**： 字符长度1-256位
	EventName string `json:"event_name"`

	AgentEventInfo *IsolateEventResponseInfo `json:"agent_event_info"`

	AntivirusResultInfo *AntivirusResultDetailInfo `json:"antivirus_result_info"`
}

func (o IsolatedFileResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsolatedFileResponseInfo struct{}"
	}

	return strings.Join([]string{"IsolatedFileResponseInfo", string(data)}, " ")
}
