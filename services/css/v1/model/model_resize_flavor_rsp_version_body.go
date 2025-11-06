package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeFlavorRspVersionBody struct {

	// **参数解释**： 引擎版本id。 **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 引擎版本名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 可变更的目标规格。 **取值范围**： 不涉及
	Flavors *[]ResizeFlavor `json:"flavors,omitempty"`
}

func (o ResizeFlavorRspVersionBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeFlavorRspVersionBody struct{}"
	}

	return strings.Join([]string{"ResizeFlavorRspVersionBody", string(data)}, " ")
}
