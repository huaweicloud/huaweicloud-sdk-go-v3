package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetPrivateZoneProxyPatternRequest Request Object
type SetPrivateZoneProxyPatternRequest struct {

	// **参数解释：** 域名ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneId string `json:"zone_id"`

	Body *SetPrivateZoneProxyPatternRequestBody `json:"body,omitempty"`
}

func (o SetPrivateZoneProxyPatternRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPrivateZoneProxyPatternRequest struct{}"
	}

	return strings.Join([]string{"SetPrivateZoneProxyPatternRequest", string(data)}, " ")
}
