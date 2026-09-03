package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingFlavorMaxAvailableResourceRequest Request Object
type ShowTrainingFlavorMaxAvailableResourceRequest struct {

	// **参数解释**：资源规格ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	FlavorId string `json:"flavor_id"`

	// **参数解释**：资源池ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolId string `json:"pool_id"`
}

func (o ShowTrainingFlavorMaxAvailableResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingFlavorMaxAvailableResourceRequest struct{}"
	}

	return strings.Join([]string{"ShowTrainingFlavorMaxAvailableResourceRequest", string(data)}, " ")
}
