package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivateZoneNameServerRequest Request Object
type ShowPrivateZoneNameServerRequest struct {

	// **参数解释：** 域名ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneId string `json:"zone_id"`
}

func (o ShowPrivateZoneNameServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateZoneNameServerRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateZoneNameServerRequest", string(data)}, " ")
}
