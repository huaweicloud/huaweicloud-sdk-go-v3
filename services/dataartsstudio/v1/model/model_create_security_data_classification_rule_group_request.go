package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityDataClassificationRuleGroupRequest Request Object
type CreateSecurityDataClassificationRuleGroupRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *DataClassificationGroupCreateDto `json:"body,omitempty"`
}

func (o CreateSecurityDataClassificationRuleGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityDataClassificationRuleGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityDataClassificationRuleGroupRequest", string(data)}, " ")
}
