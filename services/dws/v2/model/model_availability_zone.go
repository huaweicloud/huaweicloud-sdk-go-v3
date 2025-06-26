package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailabilityZone **参数解释**： 可用区详情。 **取值范围**： 不涉及。
type AvailabilityZone struct {

	// **参数解释**： 可用区唯一编码。 **取值范围**： 不涉及。
	Code string `json:"code"`

	// **参数解释**： 可用区名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 可用区状态。 **取值范围**： available：可用。 unavailable：不可用。
	Status string `json:"status"`

	// **参数解释**： 可用区组，如：center。 **取值范围**： 不涉及。
	PublicBorderGroup string `json:"public_border_group"`
}

func (o AvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZone struct{}"
	}

	return strings.Join([]string{"AvailabilityZone", string(data)}, " ")
}
