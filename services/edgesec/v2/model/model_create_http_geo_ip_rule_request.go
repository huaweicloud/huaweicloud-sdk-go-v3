package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttpGeoIpRuleRequest Request Object
type CreateHttpGeoIpRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *CreateHttpGeoIpRuleRequestBody `json:"body,omitempty"`
}

func (o CreateHttpGeoIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpGeoIpRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateHttpGeoIpRuleRequest", string(data)}, " ")
}
