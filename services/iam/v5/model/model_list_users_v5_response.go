package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersV5Response Response Object
type ListUsersV5Response struct {

	// IAM用户列表。
	Users *[]User `json:"users,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListUsersV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersV5Response struct{}"
	}

	return strings.Join([]string{"ListUsersV5Response", string(data)}, " ")
}
