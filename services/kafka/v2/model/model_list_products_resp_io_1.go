package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProductsRespIo1 struct {

	// **参数解释**： IO类型。 **取值范围**： 不涉及。
	IoType *string `json:"io_type,omitempty"`

	// **参数解释**： IO规格。 **取值范围**： 不涉及。
	StorageSpecCode *string `json:"storage_spec_code,omitempty"`

	// **参数解释**： IO未售罄的可用区列表。
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// **参数解释**： IO已售罄的不可用区列表。
	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`

	// **参数解释**： 磁盘类型。 **取值范围**： 不涉及。
	VolumeType *string `json:"volume_type,omitempty"`
}

func (o ListProductsRespIo1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespIo1 struct{}"
	}

	return strings.Join([]string{"ListProductsRespIo1", string(data)}, " ")
}
