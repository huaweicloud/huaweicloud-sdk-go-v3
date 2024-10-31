package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpPunishmentRuleResponse Response Object
type DeleteHttpPunishmentRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHttpPunishmentRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpPunishmentRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteHttpPunishmentRuleResponse", string(data)}, " ")
}
