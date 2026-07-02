package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeInstanceSpecificationResponse Response Object
type BatchChangeInstanceSpecificationResponse struct {

	// **参数解释**：  批量规格变更的任务ID列表，仅变更按需实例时会返回该参数。  **取值范围**：  不涉及。
	JobIds *[]string `json:"job_ids,omitempty"`

	// **参数解释**：  批量规格变更订单ID列表，仅变更包年/包月实例时会返回该参数。  **取值范围**：  不涉及。
	OrderIds       *[]string `json:"order_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchChangeInstanceSpecificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeInstanceSpecificationResponse struct{}"
	}

	return strings.Join([]string{"BatchChangeInstanceSpecificationResponse", string(data)}, " ")
}
