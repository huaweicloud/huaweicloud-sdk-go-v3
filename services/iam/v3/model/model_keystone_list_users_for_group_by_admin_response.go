package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type KeystoneListUsersForGroupByAdminResponse struct {
	Links *Links `json:"links,omitempty" xml:"links"`

	// IAM用户信息列表。
	Users          *[]KeystoneUserResult `json:"users,omitempty" xml:"users"`
	HttpStatusCode int                   `json:"-"`
}

func (o KeystoneListUsersForGroupByAdminResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListUsersForGroupByAdminResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListUsersForGroupByAdminResponse", string(data)}, " ")
}
