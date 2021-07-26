package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListGeoipRuleResponse struct {
	// 该策略下地理位置封禁数量

	Total *int32 `json:"total,omitempty"`
	// 地理位置封禁列表

	Items          *[]ListGeoIpResponseBodyItems `json:"items,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListGeoipRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListGeoipRuleResponse struct{}"
	}

	return strings.Join([]string{"ListGeoipRuleResponse", string(data)}, " ")
}
