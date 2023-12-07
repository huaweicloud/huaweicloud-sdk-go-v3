package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityDataClassificationRuleRequest Request Object
type DeleteSecurityDataClassificationRuleRequest struct {

	// workspace 信息
	Workspace string `json:"workspace"`

	// 需要删除的规则id
	Id string `json:"id"`
}

func (o DeleteSecurityDataClassificationRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityDataClassificationRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityDataClassificationRuleRequest", string(data)}, " ")
}
