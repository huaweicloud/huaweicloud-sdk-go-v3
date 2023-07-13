package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QaFlowHitResult struct {

	// seesionID
	SessionId *string `json:"session_id,omitempty"`

	CurrentNode *QaFlowHitNodeVo `json:"current_node,omitempty"`

	// 备用节点
	CandidateNodes *[]QaFlowHitNodeVo `json:"candidate_nodes,omitempty"`

	// 是否完整
	IsCompleted *bool `json:"is_completed,omitempty"`
}

func (o QaFlowHitResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QaFlowHitResult struct{}"
	}

	return strings.Join([]string{"QaFlowHitResult", string(data)}, " ")
}
