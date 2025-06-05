package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JarPackageHostInfo **参数解释** 服务器列表
type JarPackageHostInfo struct {

	// **参数解释** agent_id **取值范围**: 字符长度1-64位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释** 主机id **取值范围**: 字符长度0-128位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释** 服务器名称 **取值范围**: 字符长度0-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释** 服务器ip **取值范围**: 字符长度0-128位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释** 中间件包名称 **取值范围**: 字符长度0-256位
	FileName *string `json:"file_name,omitempty"`

	// **参数解释** 中间件包名称(不带后缀) **取值范围**: 字符长度0-256位\"
	Name *string `json:"name,omitempty"`

	// **参数解释** 中间件包类型 **取值范围**: 字符长度0-32位
	Catalogue *string `json:"catalogue,omitempty"`

	// **参数解释** 中间件包后缀 **取值范围**: 字符长度0-32位
	FileType *string `json:"file_type,omitempty"`

	// **参数解释** 中间件包版本 **取值范围**: 字符长度0-64位
	Version *string `json:"version,omitempty"`

	// **参数解释** 中间件包路径 **取值范围**: 字符长度0-512位
	Path *string `json:"path,omitempty"`

	// **参数解释** 中间件包hash **取值范围**: 字符长度0-512位
	Hash *string `json:"hash,omitempty"`

	// **参数解释** 中间件包大小 **取值范围**: 取值0-2147483647
	Size *int32 `json:"size,omitempty"`

	// **参数解释** uid **取值范围**: 取值0-2147483647
	Uid *int32 `json:"uid,omitempty"`

	// **参数解释** gid **取值范围**: 取值0-2147483647
	Gid *int32 `json:"gid,omitempty"`

	// **参数解释** 文件权限 **取值范围**: 字符长度0-32位
	Mode *string `json:"mode,omitempty"`

	// **参数解释** 进程id **取值范围**: 取值0-2147483647
	Pid *int32 `json:"pid,omitempty"`

	// **参数解释** 进程可执行文件路径 **取值范围**: 字符长度0-1024位
	ProcPath *string `json:"proc_path,omitempty"`

	// **参数解释** 容器实例id **取值范围**: 字符长度0-128位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释** 容器名称 **取值范围**: 字符长度0-256位
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释** 包路径 **取值范围**: 字符长度0-1024位
	PackagePath *string `json:"package_path,omitempty"`

	// **参数解释** 是否是嵌套包 **取值范围**: - 0: 不是嵌套包 - 1: 是嵌套包
	IsEmbedded *int32 `json:"is_embedded,omitempty"`

	// **参数解释** 扫描时间 **取值范围**: 取值0-4070880000000
	RecordTime *int64 `json:"record_time,omitempty"`
}

func (o JarPackageHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JarPackageHostInfo struct{}"
	}

	return strings.Join([]string{"JarPackageHostInfo", string(data)}, " ")
}
