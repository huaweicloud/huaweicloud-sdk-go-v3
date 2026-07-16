package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDevServerFlavorsRequest Request Object
type ListDevServerFlavorsRequest struct {

	// **参数解释**：服务类型。 **约束限制**：不涉及。 **取值范围**： - BMS：资源类型为裸金属服务器 - ECS：资源类型为弹性云服务器 - HPS：资源类型为超节点服务器  **默认取值**：不涉及。
	ServerType *string `json:"server_type,omitempty"`

	// **参数解释**：规格的CPU架构。 **约束限制**：不涉及。 **取值范围**： - X86：CPU架构为X86 - ARM：CPU架构为ARM **默认取值**：不涉及。
	Arch *string `json:"arch,omitempty"`

	// **参数解释**：计费模式。 **约束限制**：不涉及。 **取值范围**： - PRE_PAID：计费模式为包年/包月 - POST_PAID：计费模式为按需计费 **默认取值**：不涉及。
	ChargingMode *string `json:"charging_mode,omitempty"`
}

func (o ListDevServerFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevServerFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListDevServerFlavorsRequest", string(data)}, " ")
}
