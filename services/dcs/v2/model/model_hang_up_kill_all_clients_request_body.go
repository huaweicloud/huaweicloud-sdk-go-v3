package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HangUpKillAllClientsRequestBody struct {

	// 指定要kill全部会话的节点ID，kill_all_nodes为false时必填
	NodeId *string `json:"node_id,omitempty"`

	// true：Kill实例全部节点的会话 false: kill指定节点的全部会话
	KillAllNodes *bool `json:"kill_all_nodes,omitempty"`
}

func (o HangUpKillAllClientsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HangUpKillAllClientsRequestBody struct{}"
	}

	return strings.Join([]string{"HangUpKillAllClientsRequestBody", string(data)}, " ")
}
