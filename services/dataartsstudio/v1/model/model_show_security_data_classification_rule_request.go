package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityDataClassificationRuleRequest Request Object
type ShowSecurityDataClassificationRuleRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 需要查询的规则ID
	Id string `json:"id"`
}

func (o ShowSecurityDataClassificationRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityDataClassificationRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityDataClassificationRuleRequest", string(data)}, " ")
}
