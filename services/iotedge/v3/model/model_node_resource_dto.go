package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeResourceDto 节点资源详情
type NodeResourceDto struct {

	// 节点cpu，单位个
	Cpu *int64 `json:"cpu,omitempty"`

	// 节点内存，单位Byte
	Memory *int64 `json:"memory,omitempty"`

	// 磁盘，单位Byte
	Storage *int64 `json:"storage,omitempty"`

	// 容器pod数量
	Pods *int64 `json:"pods,omitempty"`
}

func (o NodeResourceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeResourceDto struct{}"
	}

	return strings.Join([]string{"NodeResourceDto", string(data)}, " ")
}
