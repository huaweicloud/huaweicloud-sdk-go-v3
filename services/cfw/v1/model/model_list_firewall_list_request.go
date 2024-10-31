package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFirewallListRequest Request Object
type ListFirewallListRequest struct {

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *QueryFireWallInstanceDto `json:"body,omitempty"`
}

func (o ListFirewallListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirewallListRequest struct{}"
	}

	return strings.Join([]string{"ListFirewallListRequest", string(data)}, " ")
}
