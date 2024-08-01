package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeInfo struct {

	// 节点可用区
	AvailableZone string `json:"available_zone"`

	// 子网ID
	SubnetId string `json:"subnet_id"`
}

func (o NodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeInfo struct{}"
	}

	return strings.Join([]string{"NodeInfo", string(data)}, " ")
}
