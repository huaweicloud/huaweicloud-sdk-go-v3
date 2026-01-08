package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopPoolsByUsersResponse Response Object
type ListDesktopPoolsByUsersResponse struct {

	// 用户列表。
	Users          *[]ShowDesktopPoolListByUsersInfo `json:"users,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListDesktopPoolsByUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopPoolsByUsersResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopPoolsByUsersResponse", string(data)}, " ")
}
