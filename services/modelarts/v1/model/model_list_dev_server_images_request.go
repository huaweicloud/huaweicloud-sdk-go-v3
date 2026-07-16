package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDevServerImagesRequest Request Object
type ListDevServerImagesRequest struct {

	// **参数解释**：server_type。 **约束限制**：不涉及。 **取值范围**：  - BMS  - ECS  - HPS **默认取值**：不涉及。
	ServerType *string `json:"server_type,omitempty"`

	// **参数解释**：规格名称。 **约束限制**：^.{1,128}$。 **取值范围**：不涉及。 **默认取值**：不涉及。
	FlavorName *string `json:"flavor_name,omitempty"`
}

func (o ListDevServerImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevServerImagesRequest struct{}"
	}

	return strings.Join([]string{"ListDevServerImagesRequest", string(data)}, " ")
}
