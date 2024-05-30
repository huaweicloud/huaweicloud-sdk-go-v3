package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorDto 集群节点规格
type FlavorDto struct {

	// 规格ID。
	Id *string `json:"id,omitempty"`

	// 规格名称。
	Name *string `json:"name,omitempty"`

	// 磁盘大小。
	Disk *int32 `json:"disk,omitempty"`

	// CPU大小。
	Cpu *int32 `json:"cpu,omitempty"`

	// 内存大小。
	Mem *int32 `json:"mem,omitempty"`
}

func (o FlavorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorDto struct{}"
	}

	return strings.Join([]string{"FlavorDto", string(data)}, " ")
}
