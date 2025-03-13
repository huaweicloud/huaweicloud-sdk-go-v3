package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityDataClassificationCombineRuleRequest Request Object
type UpdateSecurityDataClassificationCombineRuleRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 识别规则id
	Id string `json:"id"`

	Body *DataClassificationCombineRuleDto `json:"body,omitempty"`
}

func (o UpdateSecurityDataClassificationCombineRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityDataClassificationCombineRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityDataClassificationCombineRuleRequest", string(data)}, " ")
}
