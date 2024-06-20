package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorVolumeNodeInfo struct {

	// 节点使用存储类型
	VolumeType string `json:"volume_type"`

	// 节点使用的磁盘数量
	VolumeNum int32 `json:"volume_num"`

	// 节点去除副本后的有效容量
	Capacity int32 `json:"capacity"`

	// 节点存的单盘容量
	VolumeSize int32 `json:"volume_size"`
}

func (o FlavorVolumeNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorVolumeNodeInfo struct{}"
	}

	return strings.Join([]string{"FlavorVolumeNodeInfo", string(data)}, " ")
}
