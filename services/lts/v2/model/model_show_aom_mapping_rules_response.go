package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAomMappingRulesResponse Response Object
type ShowAomMappingRulesResponse struct {
	Body           *[]AomMappingRuleResp `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowAomMappingRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAomMappingRulesResponse struct{}"
	}

	return strings.Join([]string{"ShowAomMappingRulesResponse", string(data)}, " ")
}
