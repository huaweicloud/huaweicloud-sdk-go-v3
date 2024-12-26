package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIpsRuleModeResponse Response Object
type ChangeIpsRuleModeResponse struct {
	Data           *IpsRuleChangeRespBody `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ChangeIpsRuleModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIpsRuleModeResponse struct{}"
	}

	return strings.Join([]string{"ChangeIpsRuleModeResponse", string(data)}, " ")
}
