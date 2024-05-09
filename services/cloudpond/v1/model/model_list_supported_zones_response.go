package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportedZonesResponse Response Object
type ListSupportedZonesResponse struct {

	// 地区列表
	Zones *[]ZoneDetail `json:"zones,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSupportedZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportedZonesResponse struct{}"
	}

	return strings.Join([]string{"ListSupportedZonesResponse", string(data)}, " ")
}
