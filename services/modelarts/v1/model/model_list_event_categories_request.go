package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventCategoriesRequest Request Object
type ListEventCategoriesRequest struct {

	// **参数解释**：规格类型 **约束限制**：不涉及。 **取值范围**：可选值如下： - CPU - GPU - Ascend  **默认取值**：不涉及。
	FlavorType *string `json:"flavor_type,omitempty"`
}

func (o ListEventCategoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventCategoriesRequest struct{}"
	}

	return strings.Join([]string{"ListEventCategoriesRequest", string(data)}, " ")
}
