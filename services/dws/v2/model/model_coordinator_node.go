package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CN节点详情。
type CoordinatorNode struct {

	// 节点ID。
	Id string `json:"id"`

	// 节点名称。
	Name string `json:"name"`

	// 内网IP。
	PrivateIp string `json:"private_ip"`
}

func (o CoordinatorNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CoordinatorNode struct{}"
	}

	return strings.Join([]string{"CoordinatorNode", string(data)}, " ")
}
