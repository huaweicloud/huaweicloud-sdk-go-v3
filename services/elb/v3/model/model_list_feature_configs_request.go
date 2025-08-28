package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeatureConfigsRequest Request Object
type ListFeatureConfigsRequest struct {

	// **参数解释**：特性名称。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Feature *string `json:"feature,omitempty"`
}

func (o ListFeatureConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeatureConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListFeatureConfigsRequest", string(data)}, " ")
}
