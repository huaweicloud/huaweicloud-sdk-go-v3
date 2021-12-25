package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘节点的终端设备信息
type NodeUpdateByDevice struct {
	Added *EdgeNodeAdded `json:"added,omitempty"`
	// 要解绑的终端设备ID

	Removed *[]string `json:"removed,omitempty"`
}

func (o NodeUpdateByDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeUpdateByDevice struct{}"
	}

	return strings.Join([]string{"NodeUpdateByDevice", string(data)}, " ")
}
