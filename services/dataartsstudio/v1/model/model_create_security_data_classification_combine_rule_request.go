package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityDataClassificationCombineRuleRequest Request Object
type CreateSecurityDataClassificationCombineRuleRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *DataClassificationCombineRuleDto `json:"body,omitempty"`
}

func (o CreateSecurityDataClassificationCombineRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityDataClassificationCombineRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityDataClassificationCombineRuleRequest", string(data)}, " ")
}
