package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersResponse Response Object
type ListUsersResponse struct {

	// **参数解释**：授权对象列表。
	Users          *[]UserInfo `json:"users,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersResponse struct{}"
	}

	return strings.Join([]string{"ListUsersResponse", string(data)}, " ")
}
