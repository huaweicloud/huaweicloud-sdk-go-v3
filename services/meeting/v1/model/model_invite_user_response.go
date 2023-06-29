package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InviteUserResponse Response Object
type InviteUserResponse struct {

	// 用户是否存在
	UserExist      *bool `json:"userExist,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o InviteUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteUserResponse struct{}"
	}

	return strings.Join([]string{"InviteUserResponse", string(data)}, " ")
}
