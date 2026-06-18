package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResizeFlavorResponse Response Object
type BatchResizeFlavorResponse struct {

	// **参数解释**：  批量任务ID，仅按需实例场景返回。该任务ID不会在任务中心显示，请前往任务中心查看每个实例的变更任务。  **约束限制**：   不涉及。   **取值范围**：  不涉及。   **默认取值**：   不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：  订单ID。仅包周期实例场景返回。  **约束限制**：   不涉及。   **取值范围**：  不涉及。   **默认取值**：   不涉及。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchResizeFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResizeFlavorResponse struct{}"
	}

	return strings.Join([]string{"BatchResizeFlavorResponse", string(data)}, " ")
}
