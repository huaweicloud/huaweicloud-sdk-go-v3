package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImmutableRuleResponse Response Object
type CreateImmutableRuleResponse struct {

	// 策略ID
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateImmutableRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImmutableRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateImmutableRuleResponse", string(data)}, " ")
}
