package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImmutableRuleResponse Response Object
type DeleteImmutableRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteImmutableRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImmutableRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteImmutableRuleResponse", string(data)}, " ")
}
