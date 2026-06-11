package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailabilityZoneDetail 主备版的多AZ可用区详情
type AvailabilityZoneDetail struct {

	// **参数解释：** 主可用区，应为单可用区且和备可用区不同。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	PrimaryAvailabilityZone string `json:"primary_availability_zone"`

	// **参数解释：** 备可用区，应为单可用区且和主可用区不同。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SecondaryAvailabilityZone string `json:"secondary_availability_zone"`
}

func (o AvailabilityZoneDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZoneDetail struct{}"
	}

	return strings.Join([]string{"AvailabilityZoneDetail", string(data)}, " ")
}
