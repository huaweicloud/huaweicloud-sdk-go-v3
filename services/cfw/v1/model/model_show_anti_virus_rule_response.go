package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAntiVirusRuleResponse Response Object
type ShowAntiVirusRuleResponse struct {
	Data           *AntiVirusRuleVo `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowAntiVirusRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAntiVirusRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowAntiVirusRuleResponse", string(data)}, " ")
}
