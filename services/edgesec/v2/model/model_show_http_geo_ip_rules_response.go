package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpGeoIpRulesResponse Response Object
type ShowHttpGeoIpRulesResponse struct {

	// 策略下防护规则数量
	Total *int32 `json:"total,omitempty"`

	// 防护规则列表
	Items          *[]ShowHttpGeoIpRuleResponseBody `json:"items,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ShowHttpGeoIpRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpGeoIpRulesResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpGeoIpRulesResponse", string(data)}, " ")
}
