package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAomMappingRuleResponse struct {
	Body           *[]AomMappingRuleResp `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowAomMappingRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAomMappingRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowAomMappingRuleResponse", string(data)}, " ")
}
