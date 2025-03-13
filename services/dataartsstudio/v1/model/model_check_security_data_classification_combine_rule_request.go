package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckSecurityDataClassificationCombineRuleRequest Request Object
type CheckSecurityDataClassificationCombineRuleRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *DataClassificationCombineRuleCheckDto `json:"body,omitempty"`
}

func (o CheckSecurityDataClassificationCombineRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSecurityDataClassificationCombineRuleRequest struct{}"
	}

	return strings.Join([]string{"CheckSecurityDataClassificationCombineRuleRequest", string(data)}, " ")
}
