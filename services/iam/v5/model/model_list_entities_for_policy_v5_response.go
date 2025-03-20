package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEntitiesForPolicyV5Response Response Object
type ListEntitiesForPolicyV5Response struct {

	// 委托及信任委托列表。
	PolicyAgencies *[]PolicyAgency `json:"policy_agencies,omitempty"`

	// 用户组列表。
	PolicyGroups *[]PolicyGroup `json:"policy_groups,omitempty"`

	// IAM用户列表。
	PolicyUsers *[]PolicyUser `json:"policy_users,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListEntitiesForPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEntitiesForPolicyV5Response struct{}"
	}

	return strings.Join([]string{"ListEntitiesForPolicyV5Response", string(data)}, " ")
}
