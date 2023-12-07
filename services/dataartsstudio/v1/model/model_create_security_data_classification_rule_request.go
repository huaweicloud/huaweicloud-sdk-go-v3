package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityDataClassificationRuleRequest Request Object
type CreateSecurityDataClassificationRuleRequest struct {

	// workspace 信息
	Workspace string `json:"workspace"`

	Body *DataClassificationRuleOperateDto `json:"body,omitempty"`
}

func (o CreateSecurityDataClassificationRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityDataClassificationRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityDataClassificationRuleRequest", string(data)}, " ")
}
