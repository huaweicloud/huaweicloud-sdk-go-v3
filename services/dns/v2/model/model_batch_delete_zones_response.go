package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteZonesResponse Response Object
type BatchDeleteZonesResponse struct {

	// **参数解释：** 域名列表信息。 **取值范围：** 不涉及。
	Zones *[]ZoneData `json:"zones,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteZonesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteZonesResponse", string(data)}, " ")
}
