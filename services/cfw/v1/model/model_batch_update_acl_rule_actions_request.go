package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAclRuleActionsRequest Request Object
type BatchUpdateAclRuleActionsRequest struct {

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *UpdateSecurityPolciesActionDto `json:"body,omitempty"`
}

func (o BatchUpdateAclRuleActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAclRuleActionsRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateAclRuleActionsRequest", string(data)}, " ")
}
