package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DedicatedHostType 专属主机类型。
type DedicatedHostType struct {

	// 专属主机类型。
	HostType string `json:"host_type"`

	// 专属主机类型的vCPU数量。
	Vcpus int32 `json:"vcpus"`

	// 专属主机类型的核心数量。
	Cores int32 `json:"cores"`

	// 专属主机类型的物理套接字数量。
	Sockets int32 `json:"sockets"`

	// 专属主机类型的内存大小。
	Memory int32 `json:"memory"`

	// 专属主机规格列表。
	SupportedFlavors []string `json:"supported_flavors"`

	// 专属主机类型的类别。
	Category string `json:"category"`

	AvailabilityZoneOfferings *[]DedicatedHostTypeOffering `json:"availability_zone_offerings,omitempty"`

	PageInfo *PageInfo `json:"page_info"`
}

func (o DedicatedHostType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DedicatedHostType struct{}"
	}

	return strings.Join([]string{"DedicatedHostType", string(data)}, " ")
}
