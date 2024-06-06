package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerInfo struct {

	// 服务器ID。
	ServerId string `json:"server_id"`

	// 服务器名称。
	ServerName string `json:"server_name"`

	// ECS或者BMS。
	ServerType string `json:"server_type"`

	// 数据盘。
	DataVolumes *[]VolumeInfo `json:"data_volumes,omitempty"`

	RootVolume *VolumeInfo `json:"root_volume,omitempty"`

	// CPU类型。x86或者ARM。
	CpuType *string `json:"cpu_type,omitempty"`

	// CPU大小。
	Cpu *string `json:"cpu,omitempty"`

	// 内存大小。
	Mem *string `json:"mem,omitempty"`

	// 内部IP。
	InternalIp *string `json:"internal_ip,omitempty"`
}

func (o ServerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerInfo struct{}"
	}

	return strings.Join([]string{"ServerInfo", string(data)}, " ")
}
