package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserResponse Response Object
type ListUserResponse struct {

	// 用户数
	Count *int32 `json:"count,omitempty"`

	// 用户信息列表
	Users          *[]GetUserRsp `json:"users,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserResponse struct{}"
	}

	return strings.Join([]string{"ListUserResponse", string(data)}, " ")
}
