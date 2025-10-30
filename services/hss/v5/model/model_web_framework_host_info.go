package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebFrameworkHostInfo **参数解释** 服务器列表
type WebFrameworkHostInfo struct {

	// **参数解释** agent id **取值范围** 字符长度1-64
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释** 主机id **取值范围** 字符长度1-64
	HostId *string `json:"host_id,omitempty"`

	// **参数解释** 服务器名称 **取值范围** 字符长度1-64
	HostName *string `json:"host_name,omitempty"`

	// **参数解释** 服务器ip **取值范围** 字符长度1-64
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: web框架名称 **取值范围**: 字符长度1-256
	Name *string `json:"name,omitempty"`

	// **参数解释**: web框架版本 **取值范围**: 字符长度1-512
	Version *string `json:"version,omitempty"`

	// **参数解释**: web框架文件路径 **取值范围**: 字符长度1-512
	Path *string `json:"path,omitempty"`

	// **参数解释**: web框架扫描时间 **取值范围**: 最小值0，最大值2^63-1
	RecordTime *int64 `json:"record_time,omitempty"`

	// **参数解释**: 软件的类型 **取值范围**: 字符长度1-32
	Catalogue *string `json:"catalogue,omitempty"`

	// **参数解释**: web框架文件名称 **取值范围**: 字符长度1-256
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**: web框架文件类型 **取值范围**: 字符长度1-32
	FileType *string `json:"file_type,omitempty"`

	// **参数解释**: web框架进程gid **取值范围**: 最小值0，最大值2147483647
	Gid *int32 `json:"gid,omitempty"`

	// **参数解释**: web框架文件哈希 **取值范围**: 字符长度1-64
	Hash *string `json:"hash,omitempty"`

	// **参数解释**: 是否是压缩的文件 **取值范围**: - 0：不是压缩文件 - 1：是压缩文件
	IsEmbedded *int32 `json:"is_embedded,omitempty"`

	// **参数解释**: web框架文件权限 **取值范围**: 字符长度1-32
	Mode *string `json:"mode,omitempty"`

	// **参数解释**: web框架进程pid **取值范围**: 最小值0，最大值2147483647
	Pid *int32 `json:"pid,omitempty"`

	// **参数解释**: web框架进程路径 **取值范围**: 字符长度1-1024
	ProcPath *string `json:"proc_path,omitempty"`

	// **参数解释**: web框架进程uid **取值范围**: 最小值0，最大值2147483647
	Uid *int32 `json:"uid,omitempty"`

	// **参数解释**: 容器id **取值范围**: 字符长度1-128
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**: 容器名称 **取值范围**: 字符长度1-256
	ContainerName *string `json:"container_name,omitempty"`
}

func (o WebFrameworkHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebFrameworkHostInfo struct{}"
	}

	return strings.Join([]string{"WebFrameworkHostInfo", string(data)}, " ")
}
