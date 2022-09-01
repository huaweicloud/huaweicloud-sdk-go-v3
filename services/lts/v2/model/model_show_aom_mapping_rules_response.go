package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAomMappingRulesResponse struct {
	Body           *[]AomMappingRuleResp `json:"body,omitempty" xml:"body"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowAomMappingRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAomMappingRulesResponse struct{}"
	}

	return strings.Join([]string{"ShowAomMappingRulesResponse", string(data)}, " ")
}
