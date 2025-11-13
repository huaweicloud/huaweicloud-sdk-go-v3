package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowZoneNameServerRequest Request Object
type ShowZoneNameServerRequest struct {

	// **参数解释：** 待查询的公网域名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	DomainName string `json:"domain_name"`
}

func (o ShowZoneNameServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowZoneNameServerRequest struct{}"
	}

	return strings.Join([]string{"ShowZoneNameServerRequest", string(data)}, " ")
}
