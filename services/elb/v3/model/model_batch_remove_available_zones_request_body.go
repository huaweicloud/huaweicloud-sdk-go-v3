package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchRemoveAvailableZonesRequestBody struct {

	// **参数解释**：要移除的可用区列表。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	AvailabilityZoneList []string `json:"availability_zone_list"`
}

func (o BatchRemoveAvailableZonesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveAvailableZonesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRemoveAvailableZonesRequestBody", string(data)}, " ")
}
