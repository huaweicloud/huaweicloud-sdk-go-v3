package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPublicZonesResponse struct {
	Links *PageLink `json:"links,omitempty" xml:"links"`

	// 查询公网Zone的列表响应。
	Zones *[]PublicZoneResp `json:"zones,omitempty" xml:"zones"`

	Metadata       *Metedata `json:"metadata,omitempty" xml:"metadata"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPublicZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicZonesResponse struct{}"
	}

	return strings.Join([]string{"ListPublicZonesResponse", string(data)}, " ")
}
