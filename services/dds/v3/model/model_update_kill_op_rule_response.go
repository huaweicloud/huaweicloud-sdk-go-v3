package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKillOpRuleResponse Response Object
type UpdateKillOpRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateKillOpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKillOpRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateKillOpRuleResponse", string(data)}, " ")
}
