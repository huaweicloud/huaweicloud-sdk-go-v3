package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksInstanceInfoVolume 实例节点存储信息。
type StarRocksInstanceInfoVolume struct {

	// 实例节点存储类型。
	Type *string `json:"type,omitempty"`

	// 实例节点存储大小。
	Size *string `json:"size,omitempty"`
}

func (o StarRocksInstanceInfoVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksInstanceInfoVolume struct{}"
	}

	return strings.Join([]string{"StarRocksInstanceInfoVolume", string(data)}, " ")
}
