package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateGeoipRuleResponse struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 地理位置封禁区域

	Geoip *string `json:"geoip,omitempty"`
	// 放行或者拦截（0拦截，1放行）

	White          *int32 `json:"white,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateGeoipRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateGeoipRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateGeoipRuleResponse", string(data)}, " ")
}
