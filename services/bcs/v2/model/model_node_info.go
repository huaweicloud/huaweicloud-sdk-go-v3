package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点信息
type NodeInfo struct {
	// 节点名

	Name *string `json:"name,omitempty"`
	// 节点ID

	Id *string `json:"id,omitempty"`
}

func (o NodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeInfo struct{}"
	}

	return strings.Join([]string{"NodeInfo", string(data)}, " ")
}
