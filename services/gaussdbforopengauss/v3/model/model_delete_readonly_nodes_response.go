package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReadonlyNodesResponse Response Object
type DeleteReadonlyNodesResponse struct {

	// **参数解释**: 删除只读节点的任务ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteReadonlyNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReadonlyNodesResponse struct{}"
	}

	return strings.Join([]string{"DeleteReadonlyNodesResponse", string(data)}, " ")
}
