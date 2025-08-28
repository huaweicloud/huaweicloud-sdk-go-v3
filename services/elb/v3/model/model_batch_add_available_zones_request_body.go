package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchAddAvailableZonesRequestBody struct {

	// **参数解释**：新增的可用区列表。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	AvailabilityZoneList []string `json:"availability_zone_list"`
}

func (o BatchAddAvailableZonesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddAvailableZonesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddAvailableZonesRequestBody", string(data)}, " ")
}
