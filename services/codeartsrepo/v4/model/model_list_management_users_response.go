package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagementUsersResponse Response Object
type ListManagementUsersResponse struct {

	// 成员列表
	Body *[]ManagementUserDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListManagementUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagementUsersResponse struct{}"
	}

	return strings.Join([]string{"ListManagementUsersResponse", string(data)}, " ")
}
