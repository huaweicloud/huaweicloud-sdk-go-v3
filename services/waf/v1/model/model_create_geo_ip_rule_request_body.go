package model

import (
	"encoding/json"

	"strings"
)

type CreateGeoIpRuleRequestBody struct {
	// 地理位置封禁区域

	Geoip *string `json:"geoip,omitempty"`
	// 放行或者拦截（0拦截，1放行）

	White *int32 `json:"white,omitempty"`
	// 创建规则时间戳

	Description *string `json:"description,omitempty"`
}

func (o CreateGeoIpRuleRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateGeoIpRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGeoIpRuleRequestBody", string(data)}, " ")
}
