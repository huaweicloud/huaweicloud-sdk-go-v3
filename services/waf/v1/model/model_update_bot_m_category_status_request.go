package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBotMCategoryStatusRequest Request Object
type UpdateBotMCategoryStatusRequest struct {

	// policyId
	PolicyId string `json:"policy_id"`

	// 类别id
	CategoryId string `json:"category_id"`

	Body *UpdateBotMRuleStatusRequestBody `json:"body,omitempty"`
}

func (o UpdateBotMCategoryStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBotMCategoryStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateBotMCategoryStatusRequest", string(data)}, " ")
}
