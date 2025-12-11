package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVirsubnetCidrReservationOption struct {

	// **参数解释**： 子网预留网段的名称。 **约束限制**： 1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 子网预留网段的描述信息。 **约束限制**： 0-255个字符，不能包含“<”和“>”。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o UpdateVirsubnetCidrReservationOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirsubnetCidrReservationOption struct{}"
	}

	return strings.Join([]string{"UpdateVirsubnetCidrReservationOption", string(data)}, " ")
}
