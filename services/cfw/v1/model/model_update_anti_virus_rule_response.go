package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAntiVirusRuleResponse Response Object
type UpdateAntiVirusRuleResponse struct {
	Data           *ResponseData `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateAntiVirusRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAntiVirusRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateAntiVirusRuleResponse", string(data)}, " ")
}
