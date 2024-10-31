package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpGeoIpRuleResponse Response Object
type DeleteHttpGeoIpRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHttpGeoIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpGeoIpRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteHttpGeoIpRuleResponse", string(data)}, " ")
}
