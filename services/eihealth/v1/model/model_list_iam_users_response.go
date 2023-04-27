package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListIamUsersResponse struct {

	// 用户列表
	Users *[]IamUserDto `json:"users,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIamUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIamUsersResponse struct{}"
	}

	return strings.Join([]string{"ListIamUsersResponse", string(data)}, " ")
}
