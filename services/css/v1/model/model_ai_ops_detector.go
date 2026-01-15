package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiOpsDetector struct {

	// **参数解释**： 当前支持的检测项ID。 **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 检测项名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 检测项说明。 **取值范围**： 不涉及
	Desc *string `json:"desc,omitempty"`
}

func (o AiOpsDetector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiOpsDetector struct{}"
	}

	return strings.Join([]string{"AiOpsDetector", string(data)}, " ")
}
