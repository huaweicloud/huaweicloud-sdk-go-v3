package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceNodesResponse Response Object
type ShowInstanceNodesResponse struct {

	// **参数解释**： 后台任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Nodes *[]NodeContextEntity `json:"nodes,omitempty"`

	// **参数解释**： 总个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceNodesResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceNodesResponse", string(data)}, " ")
}
