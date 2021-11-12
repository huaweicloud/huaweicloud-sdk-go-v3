package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListGeoIpResponseBodyItems struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 地理位置封禁区域

	Geoip *string `json:"geoip,omitempty"`
	// 放行或者拦截

	White *int32 `json:"white,omitempty"`
	// 创建规则时间戳

	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o ListGeoIpResponseBodyItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeoIpResponseBodyItems struct{}"
	}

	return strings.Join([]string{"ListGeoIpResponseBodyItems", string(data)}, " ")
}
