package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAomMappingRulesResponse struct {
	Body           *[]AomMappingRuleResp `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o CreateAomMappingRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAomMappingRulesResponse struct{}"
	}

	return strings.Join([]string{"CreateAomMappingRulesResponse", string(data)}, " ")
}
