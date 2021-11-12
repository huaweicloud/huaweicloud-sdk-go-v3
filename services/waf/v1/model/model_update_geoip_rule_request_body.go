package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 地理位置封禁body
type UpdateGeoipRuleRequestBody struct {
	// 地理位置

	Geoip string `json:"geoip"`
	// 放行或者拦截(0拦截,1放行)

	White *int32 `json:"white,omitempty"`
}

func (o UpdateGeoipRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGeoipRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateGeoipRuleRequestBody", string(data)}, " ")
}
