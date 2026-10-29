package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCodeGateRequest Request Object
type UpdateCodeGateRequest struct {

	// **参数解释：** 操作者。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Operator *string `json:"operator,omitempty"`

	Body *UpdateCodeGateRequestBody `json:"body,omitempty"`
}

func (o UpdateCodeGateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCodeGateRequest struct{}"
	}

	return strings.Join([]string{"UpdateCodeGateRequest", string(data)}, " ")
}
