package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImmutableRuleResponse Response Object
type UpdateImmutableRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateImmutableRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImmutableRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateImmutableRuleResponse", string(data)}, " ")
}
