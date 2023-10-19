package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFirewallListRequest Request Object
type ListFirewallListRequest struct {

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
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
