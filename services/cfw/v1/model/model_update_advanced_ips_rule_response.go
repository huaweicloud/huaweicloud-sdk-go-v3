package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAdvancedIpsRuleResponse Response Object
type UpdateAdvancedIpsRuleResponse struct {
	Data           *ResponseData `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateAdvancedIpsRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAdvancedIpsRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateAdvancedIpsRuleResponse", string(data)}, " ")
}
