package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockTopologyResponse Response Object
type ShowDeadLockTopologyResponse struct {
	Meta *ShowDeadLockTopologyGraphRespMeta `json:"meta,omitempty"`

	// 事务节点
	Transactions *[]ShowDeadLockTopologyGraphRespTransactions `json:"transactions,omitempty"`

	// 锁节点
	Locks *[]ShowDeadLockTopologyGraphRespLocks `json:"locks,omitempty"`

	// 边，连接节点表达关系
	Edges *[]ShowDeadLockTopologyGraphRespEdges `json:"edges,omitempty"`

	// 冲突组，每条 conflicts_with 边对应一个
	ConflictGroups *[]ShowDeadLockTopologyGraphRespConflictGroups `json:"conflict_groups,omitempty"`
	HttpStatusCode int                                            `json:"-"`
}

func (o ShowDeadLockTopologyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockTopologyResponse struct{}"
	}

	return strings.Join([]string{"ShowDeadLockTopologyResponse", string(data)}, " ")
}
