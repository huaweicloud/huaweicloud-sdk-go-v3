package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CinderListAvailabilityZonesResponse struct {
	// 查询请求返回的可用分区列表，请参见• [availabilityZoneInfo参数说明](https://support.huaweicloud.com/api-evs/evs_04_2081.html#evs_04_2081__li19751007201910)。

	AvailabilityZoneInfo *[]AzInfo `json:"availabilityZoneInfo,omitempty"`
	HttpStatusCode       int       `json:"-"`
}

func (o CinderListAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"CinderListAvailabilityZonesResponse", string(data)}, " ")
}
