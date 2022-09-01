package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type NeutronListSecurityGroupsResponse struct {

	// 安全组对象列表
	SecurityGroups *[]NeutronSecurityGroup `json:"security_groups,omitempty" xml:"security_groups"`

	// 分页信息
	SecurityGroupsLinks *[]NeutronPageLink `json:"security_groups_links,omitempty" xml:"security_groups_links"`
	HttpStatusCode      int                `json:"-"`
}

func (o NeutronListSecurityGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListSecurityGroupsResponse struct{}"
	}

	return strings.Join([]string{"NeutronListSecurityGroupsResponse", string(data)}, " ")
}
