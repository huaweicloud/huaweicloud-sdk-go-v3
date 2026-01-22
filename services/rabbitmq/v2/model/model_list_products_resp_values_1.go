package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProductsRespValues1 struct {

	// **参数解释**： 规格详情。
	Detail *[]ListProductsRespDetail1 `json:"detail,omitempty"`

	// **参数解释**： 名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 资源售罄的可用区列表。
	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`

	// **参数解释**： 有可用资源的可用区列表。
	AvailableZones *[]string `json:"available_zones,omitempty"`
}

func (o ListProductsRespValues1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespValues1 struct{}"
	}

	return strings.Join([]string{"ListProductsRespValues1", string(data)}, " ")
}
