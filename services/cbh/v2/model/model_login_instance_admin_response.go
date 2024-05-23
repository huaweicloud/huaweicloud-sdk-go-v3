package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginInstanceAdminResponse Response Object
type LoginInstanceAdminResponse struct {

	// 云堡垒机登录admin链接。
	AdminUrl       *string `json:"admin_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o LoginInstanceAdminResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginInstanceAdminResponse struct{}"
	}

	return strings.Join([]string{"LoginInstanceAdminResponse", string(data)}, " ")
}
