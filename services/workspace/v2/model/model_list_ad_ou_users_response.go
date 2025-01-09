package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAdOuUsersResponse Response Object
type ListAdOuUsersResponse struct {

	// OU对象
	UserInfos *[]AdOuUserInfo `json:"user_infos,omitempty"`

	// 用户数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 可以创建的用户数量
	EnableCreateCount *int32 `json:"enable_create_count,omitempty"`
	HttpStatusCode    int    `json:"-"`
}

func (o ListAdOuUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAdOuUsersResponse struct{}"
	}

	return strings.Join([]string{"ListAdOuUsersResponse", string(data)}, " ")
}
