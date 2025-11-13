package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInteractionRuleResponse Response Object
type DeleteInteractionRuleResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInteractionRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInteractionRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteInteractionRuleResponse", string(data)}, " ")
}
