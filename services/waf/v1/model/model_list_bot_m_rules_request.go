package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBotMRulesRequest Request Object
type ListBotMRulesRequest struct {

	// policyid
	PolicyId string `json:"policy_id"`
}

func (o ListBotMRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBotMRulesRequest struct{}"
	}

	return strings.Join([]string{"ListBotMRulesRequest", string(data)}, " ")
}
