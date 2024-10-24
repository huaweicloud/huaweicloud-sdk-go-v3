package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnlargeNodeInfo struct {

	// 节点可用区
	AvailableZone string `json:"available_zone"`
}

func (o EnlargeNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeNodeInfo struct{}"
	}

	return strings.Join([]string{"EnlargeNodeInfo", string(data)}, " ")
}
