package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginCbhResponse Response Object
type LoginCbhResponse struct {

	// IAM登录链接
	LoginUrl       *string `json:"login_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o LoginCbhResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginCbhResponse struct{}"
	}

	return strings.Join([]string{"LoginCbhResponse", string(data)}, " ")
}
