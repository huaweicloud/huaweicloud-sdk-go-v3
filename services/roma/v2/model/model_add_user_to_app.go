package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddUserToApp struct {
	// 要添加的用户成员列表，空列表时代表清空应用的所有成员

	Users *[]User `json:"users,omitempty"`
}

func (o AddUserToApp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddUserToApp struct{}"
	}

	return strings.Join([]string{"AddUserToApp", string(data)}, " ")
}
