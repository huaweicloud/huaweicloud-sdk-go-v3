package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResetPoolNodesResponse Response Object
type BatchResetPoolNodesResponse struct {

	// **参数解释**：异步任务的ID。 **取值范围**：不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchResetPoolNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResetPoolNodesResponse struct{}"
	}

	return strings.Join([]string{"BatchResetPoolNodesResponse", string(data)}, " ")
}
