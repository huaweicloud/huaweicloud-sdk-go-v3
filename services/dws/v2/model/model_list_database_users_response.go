package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseUsersResponse Response Object
type ListDatabaseUsersResponse struct {

	// **参数解释**： 用户/角色列表。 **取值范围**： 不涉及。
	Users          *[]DatabaseUser `json:"users,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListDatabaseUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseUsersResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseUsersResponse", string(data)}, " ")
}
