package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebAppAndServiceResponseInfo struct {

	// **参数解释**: web应用、web服务、数据库资产指纹种类 **取值范围**: 字符长度0-64
	Catalogue *string `json:"catalogue,omitempty"`

	// **参数解释**: 资产指纹名字 **取值范围**: 字符长度0-64
	Name *string `json:"name,omitempty"`

	// **参数解释**: 资产指纹版本 **取值范围**: 字符长度0-64
	Version *string `json:"version,omitempty"`

	// **参数解释** agent id **取值范围** 字符长度1-64
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释** 安装路径 **取值范围** 字符长度0-256
	InstallPath *string `json:"install_path,omitempty"`

	// **参数解释** 配置文件路径 **取值范围** 字符长度0-256
	ConfigPath *string `json:"config_path,omitempty"`

	// **参数解释**: uid **取值范围**: 最小值0，最大值2147483647
	Uid *int32 `json:"uid,omitempty"`

	// **参数解释**: gid **取值范围**: 最小值0，最大值2147483647
	Gid *int32 `json:"gid,omitempty"`

	// **参数解释**: 资产指纹文件权限 **取值范围**: 字符长度1-32
	Mode *string `json:"mode,omitempty"`

	// **参数解释**: 资产指纹文件最近状态改变时间 **取值范围**: 最小值0，最大值2^63-1
	Ctime *int64 `json:"ctime,omitempty"`

	// **参数解释**: 资产指纹文件最近修改时间 **取值范围**: 最小值0，最大值2^63-1
	Mtime *int64 `json:"mtime,omitempty"`

	// **参数解释**: 资产指纹文件最近访问时间 **取值范围**: 最小值0，最大值2^63-1
	Atime *int64 `json:"atime,omitempty"`

	// **参数解释**: pid **取值范围**: 最小值0，最大值2147483647
	Pid *int32 `json:"pid,omitempty"`

	// **参数解释**: 资产指纹进程路径 **取值范围**: 字符长度0-256
	ProcPath *string `json:"proc_path,omitempty"`

	// **参数解释**: 容器id **取值范围**: 字符长度0-256
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**: 容器名称 **取值范围**: 字符长度0-256
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**: 资产指纹扫描时间 **取值范围**: 最小值0，最大值2^63-1
	RecordTime *int64 `json:"record_time,omitempty"`

	// **参数解释** 主机id **取值范围** 字符长度1-64
	HostId *string `json:"host_id,omitempty"`

	// **参数解释** 服务器名称 **取值范围** 字符长度1-64
	HostName *string `json:"host_name,omitempty"`

	// **参数解释** 服务器ip **取值范围** 字符长度1-64
	HostIp *string `json:"host_ip,omitempty"`
}

func (o WebAppAndServiceResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebAppAndServiceResponseInfo struct{}"
	}

	return strings.Join([]string{"WebAppAndServiceResponseInfo", string(data)}, " ")
}
