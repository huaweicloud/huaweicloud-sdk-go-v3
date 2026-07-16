package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerScaleEvaluation struct {

	// **参数解释**：是否售罄。 **约束限制**：不涉及 **取值范围**：不涉及。 **默认取值**：不涉及。
	IsSoldOut *bool `json:"is_sold_out,omitempty"`

	// **参数解释**：规格信息。 **约束限制**：不涉及 **取值范围**：不涉及。 **默认取值**：不涉及。
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**：资源规格信息。 **约束限制**：不涉及 **取值范围**：不涉及。 **默认取值**：不涉及。
	ResourceFlavor *string `json:"resource_flavor,omitempty"`
}

func (o ServerScaleEvaluation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerScaleEvaluation struct{}"
	}

	return strings.Join([]string{"ServerScaleEvaluation", string(data)}, " ")
}
