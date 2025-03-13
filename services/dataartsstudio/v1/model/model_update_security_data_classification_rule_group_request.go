package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityDataClassificationRuleGroupRequest Request Object
type UpdateSecurityDataClassificationRuleGroupRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 需要查询的规则组ID
	Id string `json:"id"`

	Body *DataClassificationGroupUpdateDto `json:"body,omitempty"`
}

func (o UpdateSecurityDataClassificationRuleGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityDataClassificationRuleGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityDataClassificationRuleGroupRequest", string(data)}, " ")
}
