package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAuthorizedDbUsersResponse struct {

	// 用户及相关权限。
	Users *[]UserWithPrivilege `json:"users,omitempty" xml:"users"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuthorizedDbUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizedDbUsersResponse struct{}"
	}

	return strings.Join([]string{"ListAuthorizedDbUsersResponse", string(data)}, " ")
}
