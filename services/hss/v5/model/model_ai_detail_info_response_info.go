package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiDetailInfoResponseInfo **参数解释** AI组件详情信息
type AiDetailInfoResponseInfo struct {

	// **参数解释**： category==host时 表示服务器名称 category==container时 表示节点名称 category==serverless时 表示实例名称 **取值范围**： 字符长度1-256位
	ServerName *string `json:"server_name,omitempty"`

	// **参数解释**： category==host时 表示服务器IP地址 category==container时 表示节点IP地址 category==serverless时 表示实例IP地址        **取值范围**： 字符长度1-128位
	ServerIp *string `json:"server_ip,omitempty"`

	// **参数解释**： AI应用名称 **取值范围**： 字符长度1-256位
	AiApplication *string `json:"ai_application,omitempty"`

	// **参数解释**： AI工具名称 **取值范围**： 字符长度1-256位
	AiTool *string `json:"ai_tool,omitempty"`

	// **参数解释**： AI应用类型 **取值范围**： 字符长度1-256位
	Type *string `json:"type,omitempty"`

	// **参数解释**： 版本号 **取值范围**： 字符长度1-64位
	Version *string `json:"version,omitempty"`

	// **参数解释**： 应用启动路径 **取值范围**： 字符长度1-512位
	StartupPath *string `json:"startup_path,omitempty"`

	// **参数解释**： 应用启动时间 **取值范围**： 时间戳，毫秒级
	StartupTime *int64 `json:"startup_time,omitempty"`

	// **参数解释**： 安装路径 **取值范围**： 字符长度1-512位
	InstallPath *string `json:"install_path,omitempty"`

	// **参数解释**： 应用启动命令行 **取值范围**： 字符长度1-512位
	Cmdline *string `json:"cmdline,omitempty"`

	// **参数解释**： 首次扫描时间 **取值范围**： 时间戳，毫秒级
	FirstScanTime *int64 `json:"first_scan_time,omitempty"`

	// **参数解释**： 最近一次扫描时间 **取值范围**： 时间戳，毫秒级
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释**： 容器名称 **取值范围**： 字符长度1-256位
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**： 容器ID **取值范围**： 字符长度1-128位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 服务器ID，查看服务器详情使用 **取值范围**： 字符长度1-128位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**： 应用进程PID **取值范围**： 取值0-2147483647
	Pid *int32 `json:"pid,omitempty"`

	// **参数解释**： 应用进程的父进程PID **取值范围**： 取值0-2147483647
	Ppid *int32 `json:"ppid,omitempty"`

	// **参数解释**： 应用运行用户 **取值范围**： 字符长度1-128位
	User *string `json:"user,omitempty"`

	// **参数解释**： 应用进程监听网络信息
	NetInfo *[]AiProcessNetInfo `json:"net_info,omitempty"`
}

func (o AiDetailInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiDetailInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"AiDetailInfoResponseInfo", string(data)}, " ")
}
