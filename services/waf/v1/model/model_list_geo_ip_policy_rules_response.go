package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGeoIpPolicyRulesResponse Response Object
type ListGeoIpPolicyRulesResponse struct {

	// 该策略下地理位置控制规则数量
	Total *int32 `json:"total,omitempty"`

	// 地理位置控制规则数组
	Items          *[]GeOIpItem `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListGeoIpPolicyRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeoIpPolicyRulesResponse struct{}"
	}

	return strings.Join([]string{"ListGeoIpPolicyRulesResponse", string(data)}, " ")
}
