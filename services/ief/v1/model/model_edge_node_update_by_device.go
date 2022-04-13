package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘节点的终端设备信息
type EdgeNodeUpdateByDevice struct {
	Nodes *NodeUpdateByDevice `json:"nodes"`
}

func (o EdgeNodeUpdateByDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeNodeUpdateByDevice struct{}"
	}

	return strings.Join([]string{"EdgeNodeUpdateByDevice", string(data)}, " ")
}
