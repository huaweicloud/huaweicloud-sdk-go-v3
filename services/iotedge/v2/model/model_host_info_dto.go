package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostInfoDto 主机信息
type HostInfoDto struct {

	// 主机标签DEFAULT|MASTER|SLAVE
	HostTag *string `json:"host_tag,omitempty"`

	// 主机工作状态(ONLINE|OFFLINE)
	HostStatus *string `json:"host_status,omitempty"`

	// 边缘节点操作系统。例如：Ubuntu 20.04；CentOS 7.9。不同于os_type边缘节点系统类型。
	OsName *string `json:"os_name,omitempty"`

	// 边缘节点主机名
	HostName *string `json:"host_name,omitempty"`

	// 容器运行时版本
	ContainerVersion *string `json:"container_version,omitempty"`

	// 边缘节点网络网卡信息
	Nics *[]Nic `json:"nics,omitempty"`

	// 网络规格，如4 cores | 3867 MB
	Specification *string `json:"specification,omitempty"`

	// NPU设备详细信息，包括硬件信息和使用情况。
	NpuDetails *[]NpuDetailsDto `json:"npu_details,omitempty"`
}

func (o HostInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostInfoDto struct{}"
	}

	return strings.Join([]string{"HostInfoDto", string(data)}, " ")
}
