package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIpsRuleModeRequest Request Object
type ChangeIpsRuleModeRequest struct {

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *IpsRuleChangeDto `json:"body,omitempty"`
}

func (o ChangeIpsRuleModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIpsRuleModeRequest struct{}"
	}

	return strings.Join([]string{"ChangeIpsRuleModeRequest", string(data)}, " ")
}
