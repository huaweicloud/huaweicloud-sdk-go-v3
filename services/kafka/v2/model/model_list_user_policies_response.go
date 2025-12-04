package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserPoliciesResponse Response Object
type ListUserPoliciesResponse struct {

	// **参数解释**： 用户名。 **取值范围**： 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 用户权限总数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 用户权限列表。
	Policies       *[]UserPolicyResponeEntity `json:"policies,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListUserPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListUserPoliciesResponse", string(data)}, " ")
}
