package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiskType **参数解释**： 磁盘类型。 **取值范围**： 不涉及
type DiskType struct {

	// **参数解释**： 可用区。 **取值范围**： 不涉及
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// **参数解释**： 磁盘类型。 **取值范围**： - SATA：普通IO - SAS：高IO - SSD：超高IO - ESSD：极速型SSD
	VolumeNames *[]string `json:"volumeNames,omitempty"`
}

func (o DiskType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskType struct{}"
	}

	return strings.Join([]string{"DiskType", string(data)}, " ")
}
