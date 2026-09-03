package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelConnectionProcessRequestBody Kill进程请求体
type CancelConnectionProcessRequestBody struct {

	// 查杀会话的ID列表
	ProcessIds []string `json:"process_ids"`

	// 是否查杀全部会话
	KillAll *bool `json:"kill_all,omitempty"`

	// 实例节点ID，实例节点的唯一标识
	NodeId *string `json:"node_id,omitempty"`

	// 实例节点类型（master：主节点，slave：副节点，readreplica：只读节点）
	NodeRole *string `json:"node_role,omitempty"`
}

func (o CancelConnectionProcessRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelConnectionProcessRequestBody struct{}"
	}

	return strings.Join([]string{"CancelConnectionProcessRequestBody", string(data)}, " ")
}
