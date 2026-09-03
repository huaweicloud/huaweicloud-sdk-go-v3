package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupUsageExceededInstancesResponse Response Object
type ShowBackupUsageExceededInstancesResponse struct {

	// **参数解释**：  超阈值实例列表。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Instances *[]ExceededInstanceInfo `json:"instances,omitempty"`

	// **参数解释**：  超阈值实例总数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowBackupUsageExceededInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupUsageExceededInstancesResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupUsageExceededInstancesResponse", string(data)}, " ")
}
