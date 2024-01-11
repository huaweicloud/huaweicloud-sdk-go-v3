package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityDataClassificationRuleRequest Request Object
type CreateSecurityDataClassificationRuleRequest struct {

	// DataArts Studio工作空间ID
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
