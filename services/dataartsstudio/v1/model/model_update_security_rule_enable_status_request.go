package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityRuleEnableStatusRequest Request Object
type UpdateSecurityRuleEnableStatusRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 识别规则id
	Id string `json:"id"`

	Body *DataClassificationRuleEnableDto `json:"body,omitempty"`
}

func (o UpdateSecurityRuleEnableStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityRuleEnableStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityRuleEnableStatusRequest", string(data)}, " ")
}
