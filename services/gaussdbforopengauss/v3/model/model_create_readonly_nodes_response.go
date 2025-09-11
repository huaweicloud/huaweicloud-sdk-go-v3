package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReadonlyNodesResponse Response Object
type CreateReadonlyNodesResponse struct {

	// **参数解释**: 创建只读节点的任务ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateReadonlyNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReadonlyNodesResponse struct{}"
	}

	return strings.Join([]string{"CreateReadonlyNodesResponse", string(data)}, " ")
}
