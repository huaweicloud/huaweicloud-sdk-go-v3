package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeoipRuleResponse struct{}"
	}

	return strings.Join([]string{"ListGeoipRuleResponse", string(data)}, " ")
}
