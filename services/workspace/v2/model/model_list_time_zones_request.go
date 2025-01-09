package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTimeZonesRequest Request Object
type ListTimeZonesRequest struct {

	// 按照时区name过滤
	TimeZoneName *string `json:"time_zone_name,omitempty"`
}

func (o ListTimeZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTimeZonesRequest struct{}"
	}

	return strings.Join([]string{"ListTimeZonesRequest", string(data)}, " ")
}
