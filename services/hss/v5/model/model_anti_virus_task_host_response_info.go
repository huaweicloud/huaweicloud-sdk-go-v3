package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusTaskHostResponseInfo 扫描任务关联主机信息
type AntiVirusTaskHostResponseInfo struct {

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-256位，支持IPv4或IPv6格式（IPv4长度7-15位，IPv6长度15-39位）
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 资产重要性。 **取值范围**： - important ：重要资产。 - common ：一般资产。 - test ：测试资产。
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**： 启动时间 **取值范围**： 最小值0，最大值9223372036854775807；时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算）；单位：ms
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**: 运行时长 **取值范围**: 非负整数，最小值0；单位：s（秒）
	RunDuration *int64 `json:"run_duration,omitempty"`

	// **参数解释**： 扫描进度 **取值范围**： 字符串格式，支持百分比（如“50%”）或0-100的数值字符串
	ScanProgress *string `json:"scan_progress,omitempty"`

	// **参数解释** 新发现病毒数量 **取值范围** 非负整数，最小值0；单位：个
	VirusNum *int32 `json:"virus_num,omitempty"`

	// **参数解释**: 已扫描的文件数量 **取值范围**: 非负整数，最小值0；单位：个
	ScanFileNum *int64 `json:"scan_file_num,omitempty"`

	// **参数解释**: 服务器扫描状态 **取值范围**: - scanning ：扫描中 - success ：扫描成功 - fail ：扫描失败 - cancel ：取消扫描
	HostTaskStatus *string `json:"host_task_status,omitempty"`

	// **参数解释**: 失败原因 **取值范围**: 字符长度0-512位
	FailReason *string `json:"fail_reason,omitempty"`

	// **参数解释**： 是否删除 **取值范围**： 包含如下:   - true ：已删除   - false : 未删除
	Deleted *bool `json:"deleted,omitempty"`

	// **参数解释**： 是否使用病毒查杀按次计费配额 **取值范围**： 0（未使用）、1（已使用）
	WhetherUsingQuota *int32 `json:"whether_using_quota,omitempty"`

	// **参数解释**: 主机上安装的杀毒Agent的唯一标识ID，用于关联主机与杀毒服务 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux - Windows：Windows
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 主机状态 **取值范围**: - ACTIVE：正在运行 - SHUTOFF：关机 - BUILDING：创建中 - ERROR：故障
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**： Agent状态 **取值范围**: - installed：已安装 - not_installed：未安 - online：在线 - offline：离线 - install_failed：安装失败 - installing：安装中 - not_online：不在线的（除了在线以外的所有状态，仅作为查询条件）
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: 防护状态 **取值范围**:  - closed ：关闭  - opened ：开启
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**: 操作系统名称 **取值范围**: 字符长度0-128位
	OsName *string `json:"os_name,omitempty"`

	// **参数解释**： 系统版本号 **取值范围**： 字符长度0-64位
	OsVersion *string `json:"os_version,omitempty"`
}

func (o AntiVirusTaskHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusTaskHostResponseInfo struct{}"
	}

	return strings.Join([]string{"AntiVirusTaskHostResponseInfo", string(data)}, " ")
}
