package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusTaskHostResponseInfo 扫描任务关联主机信息
type AntiVirusTaskHostResponseInfo struct {

	// **参数解释**: 服务器ID **取值范围**: 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 弹性公网IP地址
	PublicIp *string `json:"public_ip,omitempty"`

	// 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// 启动时间，毫秒
	StartTime *int64 `json:"start_time,omitempty"`

	// 运行时长，秒
	RunDuration *int64 `json:"run_duration,omitempty"`

	// 扫描进度
	ScanProgress *string `json:"scan_progress,omitempty"`

	// 新发现病毒数量
	VirusNum *int32 `json:"virus_num,omitempty"`

	// 已扫描的文件数量
	ScanFileNum *int64 `json:"scan_file_num,omitempty"`

	// 服务器扫描状态，包含如下4种   - scanning ：扫描中   - success ：扫描成功   - fail ：扫描失败   - cancel ：取消扫描
	HostTaskStatus *string `json:"host_task_status,omitempty"`

	// 失败原因
	FailReason *string `json:"fail_reason,omitempty"`

	// 是否删除，包含如下:   - true ：已删除   - false : 未删除
	Deleted *bool `json:"deleted,omitempty"`

	// 是否使用病毒查杀按次计费配额
	WhetherUsingQuota *int32 `json:"whether_using_quota,omitempty"`

	// **参数解释**: Agent ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// 操作系统类型，包含如下2种。   - Linux ：Linux。   - Windows ：Windows。
	OsType *string `json:"os_type,omitempty"`

	// 服务器状态
	HostStatus *string `json:"host_status,omitempty"`

	// Agent状态，包含如下6种。   - installed ：已安装。   - not_installed ：未安装。   - online ：在线。   - offline ：离线。   - install_failed ：安装失败。   - installing ：安装中。   - not_online ：不在线的（除了在线以外的所有状态，仅作为查询条件）。
	AgentStatus *string `json:"agent_status,omitempty"`

	// 防护状态，包含如下2种。   - closed ：关闭。   - opened ：开启。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 操作系统名称
	OsName *string `json:"os_name,omitempty"`

	// 操作系统版本
	OsVersion *string `json:"os_version,omitempty"`
}

func (o AntiVirusTaskHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusTaskHostResponseInfo struct{}"
	}

	return strings.Join([]string{"AntiVirusTaskHostResponseInfo", string(data)}, " ")
}
