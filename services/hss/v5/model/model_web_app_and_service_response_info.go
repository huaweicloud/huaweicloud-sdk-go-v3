package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebAppAndServiceResponseInfo struct {

	// 资产指纹种类
	Catalogue *string `json:"catalogue,omitempty"`

	// 资产指纹名字
	Name *string `json:"name,omitempty"`

	// 资产指纹-数据库-版本
	Version *string `json:"version,omitempty"`

	// agent_id
	AgentId *string `json:"agent_id,omitempty"`

	// install_path
	InstallPath *string `json:"install_path,omitempty"`

	// config_path 配置文件路径
	ConfigPath *string `json:"config_path,omitempty"`

	// uid
	Uid *int32 `json:"uid,omitempty"`

	// gid
	Gid *int32 `json:"gid,omitempty"`

	// mode
	Mode *string `json:"mode,omitempty"`

	// ctime
	Ctime *int64 `json:"ctime,omitempty"`

	// mtime
	Mtime *int64 `json:"mtime,omitempty"`

	// atime
	Atime *int64 `json:"atime,omitempty"`

	// pid
	Pid *int32 `json:"pid,omitempty"`

	// proc_path
	ProcPath *string `json:"proc_path,omitempty"`

	// container_id
	ContainerId *string `json:"container_id,omitempty"`

	// container_name
	ContainerName *string `json:"container_name,omitempty"`

	// record_time
	RecordTime *int64 `json:"record_time,omitempty"`

	// host_name
	HostName *string `json:"host_name,omitempty"`

	// host_id
	HostId *string `json:"host_id,omitempty"`

	// host_ip
	HostIp *string `json:"host_ip,omitempty"`
}

func (o WebAppAndServiceResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebAppAndServiceResponseInfo struct{}"
	}

	return strings.Join([]string{"WebAppAndServiceResponseInfo", string(data)}, " ")
}
