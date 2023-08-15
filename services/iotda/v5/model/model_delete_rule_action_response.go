package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRuleActionResponse Response Object
type DeleteRuleActionResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRuleActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleActionResponse struct{}"
	}

	return strings.Join([]string{"DeleteRuleActionResponse", string(data)}, " ")
}
