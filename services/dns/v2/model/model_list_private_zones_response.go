package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateZonesResponse Response Object
type ListPrivateZonesResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	// **参数解释：** 域名列表信息。 **取值范围：** 不涉及。
	Zones          *[]PrivateZoneResp `json:"zones,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListPrivateZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateZonesResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateZonesResponse", string(data)}, " ")
}
