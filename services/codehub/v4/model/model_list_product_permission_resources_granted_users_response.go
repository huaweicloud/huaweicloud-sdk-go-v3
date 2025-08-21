package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProductPermissionResourcesGrantedUsersResponse Response Object
type ListProductPermissionResourcesGrantedUsersResponse struct {

	// **参数解释：** 成员列表。
	Body           *[]GrantedUsersDto `json:"body,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListProductPermissionResourcesGrantedUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductPermissionResourcesGrantedUsersResponse struct{}"
	}

	return strings.Join([]string{"ListProductPermissionResourcesGrantedUsersResponse", string(data)}, " ")
}
