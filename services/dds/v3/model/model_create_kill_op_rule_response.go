package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKillOpRuleResponse Response Object
type CreateKillOpRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateKillOpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKillOpRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateKillOpRuleResponse", string(data)}, " ")
}
