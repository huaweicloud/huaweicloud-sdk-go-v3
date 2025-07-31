package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePublicZoneRequest Request Object
type DeletePublicZoneRequest struct {

	// **参数解释：** 域名ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneId string `json:"zone_id"`
}

func (o DeletePublicZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicZoneRequest struct{}"
	}

	return strings.Join([]string{"DeletePublicZoneRequest", string(data)}, " ")
}
