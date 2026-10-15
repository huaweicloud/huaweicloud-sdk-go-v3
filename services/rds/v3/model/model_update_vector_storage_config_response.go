package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVectorStorageConfigResponse Response Object
type UpdateVectorStorageConfigResponse struct {

	// **参数解释**：  任务ID，可用于查询异步任务执行状态。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateVectorStorageConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVectorStorageConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateVectorStorageConfigResponse", string(data)}, " ")
}
