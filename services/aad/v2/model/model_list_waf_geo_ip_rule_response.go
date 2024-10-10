package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWafGeoIpRuleResponse Response Object
type ListWafGeoIpRuleResponse struct {

	// total
	Total *int32 `json:"total,omitempty"`

	// items
	Items          *[]WafGeoIpRule `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListWafGeoIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWafGeoIpRuleResponse struct{}"
	}

	return strings.Join([]string{"ListWafGeoIpRuleResponse", string(data)}, " ")
}
