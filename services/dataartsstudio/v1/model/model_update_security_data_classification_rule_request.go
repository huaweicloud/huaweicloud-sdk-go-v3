package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityDataClassificationRuleRequest Request Object
type UpdateSecurityDataClassificationRuleRequest struct {

	// workspace 信息
	Workspace string `json:"workspace"`

	// 识别规则id
	Id string `json:"id"`

	Body *DataClassificationRuleOperateDto `json:"body,omitempty"`
}

func (o UpdateSecurityDataClassificationRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityDataClassificationRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityDataClassificationRuleRequest", string(data)}, " ")
}
