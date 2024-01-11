package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityDataClassificationRuleGroupRequest Request Object
type ShowSecurityDataClassificationRuleGroupRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 需要查询的规则组ID
	Id string `json:"id"`
}

func (o ShowSecurityDataClassificationRuleGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityDataClassificationRuleGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityDataClassificationRuleGroupRequest", string(data)}, " ")
}
