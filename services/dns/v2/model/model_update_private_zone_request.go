package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivateZoneRequest Request Object
type UpdatePrivateZoneRequest struct {

	// **参数解释：** 域名ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneId string `json:"zone_id"`

	Body *UpdatePrivateZoneInfoReq `json:"body,omitempty"`
}

func (o UpdatePrivateZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateZoneRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateZoneRequest", string(data)}, " ")
}
