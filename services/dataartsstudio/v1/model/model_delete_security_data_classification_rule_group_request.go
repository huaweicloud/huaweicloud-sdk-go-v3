package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityDataClassificationRuleGroupRequest Request Object
type DeleteSecurityDataClassificationRuleGroupRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BatchDeleteRuleGroupsBaseDto `json:"body,omitempty"`
}

func (o DeleteSecurityDataClassificationRuleGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityDataClassificationRuleGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityDataClassificationRuleGroupRequest", string(data)}, " ")
}
