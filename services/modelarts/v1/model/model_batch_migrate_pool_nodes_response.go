package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchMigratePoolNodesResponse Response Object
type BatchMigratePoolNodesResponse struct {

	// **参数解释**：异步任务的ID。 **取值范围**：不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchMigratePoolNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMigratePoolNodesResponse struct{}"
	}

	return strings.Join([]string{"BatchMigratePoolNodesResponse", string(data)}, " ")
}
