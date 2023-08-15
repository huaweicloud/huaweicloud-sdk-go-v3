package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersResponse Response Object
type ListUsersResponse struct {

	// 用户总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 用户列表。
	Users          *[]User `json:"users,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersResponse struct{}"
	}

	return strings.Join([]string{"ListUsersResponse", string(data)}, " ")
}
