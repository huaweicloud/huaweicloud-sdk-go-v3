package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPrivateZonesResponse struct {
	Links *PageLink `json:"links,omitempty" xml:"links"`

	Metadata *Metedata `json:"metadata,omitempty" xml:"metadata"`

	Zones          *[]PrivateZoneResp `json:"zones,omitempty" xml:"zones"`
	HttpStatusCode int                `json:"-"`
}

func (o ListPrivateZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateZonesResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateZonesResponse", string(data)}, " ")
}
