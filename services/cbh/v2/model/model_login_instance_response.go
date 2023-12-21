package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginInstanceResponse Response Object
type LoginInstanceResponse struct {

	// 云堡垒机登录链接。
	LoginUrl       *string `json:"login_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o LoginInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginInstanceResponse struct{}"
	}

	return strings.Join([]string{"LoginInstanceResponse", string(data)}, " ")
}
