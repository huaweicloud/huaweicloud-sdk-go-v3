package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTimeZonesResponse Response Object
type ListTimeZonesResponse struct {

	// 时区列表。
	TimeZones      *[]Timezones `json:"time_zones,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListTimeZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTimeZonesResponse struct{}"
	}

	return strings.Join([]string{"ListTimeZonesResponse", string(data)}, " ")
}
