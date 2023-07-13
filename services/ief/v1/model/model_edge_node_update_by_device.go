package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EdgeNodeUpdateByDevice 边缘节点的终端设备信息
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
