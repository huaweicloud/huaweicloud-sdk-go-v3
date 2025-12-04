package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartNodeRequest Request Object
type RestartNodeRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`

	// 节点 ID。
	NodeId string `json:"node_id"`
}

func (o RestartNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartNodeRequest struct{}"
	}

	return strings.Join([]string{"RestartNodeRequest", string(data)}, " ")
}
