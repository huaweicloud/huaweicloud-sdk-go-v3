package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttpPunishmentRuleRequest Request Object
type CreateHttpPunishmentRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *CreateHttpPunishmentRuleRequestBody `json:"body,omitempty"`
}

func (o CreateHttpPunishmentRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpPunishmentRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateHttpPunishmentRuleRequest", string(data)}, " ")
}
