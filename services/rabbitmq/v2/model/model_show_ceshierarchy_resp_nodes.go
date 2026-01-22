package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespNodes struct {

	// **参数解释**： 节点名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 可用区。 **取值范围**： 不涉及。
	AvailableZone *string `json:"available_zone,omitempty"`
}

func (o ShowCeshierarchyRespNodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespNodes struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespNodes", string(data)}, " ")
}
