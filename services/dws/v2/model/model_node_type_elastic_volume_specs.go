package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeTypeElasticVolumeSpecs 如果规格为弹性容量规格，则该属性为规格典配的弹性容量信息，包括存储类型、最小容量、最大容量以及步长信息，如果为固定存储规格，则该属性为null。
type NodeTypeElasticVolumeSpecs struct {

	// 云盘存储类型。
	Type string `json:"type"`

	// 云盘容量调整步长。
	Step string `json:"step"`

	// 云盘支持的最小容量。
	MinSize int32 `json:"min_size"`

	// 云盘支持的最大容量。
	MaxSize int32 `json:"max_size"`
}

func (o NodeTypeElasticVolumeSpecs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTypeElasticVolumeSpecs struct{}"
	}

	return strings.Join([]string{"NodeTypeElasticVolumeSpecs", string(data)}, " ")
}
