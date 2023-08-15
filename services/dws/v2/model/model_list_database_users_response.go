package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseUsersResponse Response Object
type ListDatabaseUsersResponse struct {

	// user list
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
