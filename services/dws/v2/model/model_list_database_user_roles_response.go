package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseUserRolesResponse Response Object
type ListDatabaseUserRolesResponse struct {

	// **参数解释**： 授予的权限。 **取值范围**： 不涉及。
	Roles *[]RoleMember `json:"roles,omitempty"`

	// **参数解释**： 总条数。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabaseUserRolesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseUserRolesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseUserRolesResponse", string(data)}, " ")
}
