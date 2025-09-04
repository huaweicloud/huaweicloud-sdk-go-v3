package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProxyFlavorGroup struct {

	// **参数解释**：  规格分组类型。  **取值范围**：  不涉及。
	GroupType *string `json:"group_type,omitempty"`

	// **参数解释**：  规格列表。
	ProxyFlavors *[]ProxyFlavor `json:"proxy_flavors,omitempty"`
}

func (o ProxyFlavorGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyFlavorGroup struct{}"
	}

	return strings.Join([]string{"ProxyFlavorGroup", string(data)}, " ")
}
