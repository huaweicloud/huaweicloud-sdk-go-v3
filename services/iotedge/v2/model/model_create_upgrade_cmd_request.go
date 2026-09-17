package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUpgradeCmdRequest Request Object
type CreateUpgradeCmdRequest struct {

	// 节点id
	EdgeNodeId string `json:"edge_node_id"`
}

func (o CreateUpgradeCmdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUpgradeCmdRequest struct{}"
	}

	return strings.Join([]string{"CreateUpgradeCmdRequest", string(data)}, " ")
}
