package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetPwdModeResponse Response Object
type ResetPwdModeResponse struct {

	// 密码
	Password       *string `json:"password,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetPwdModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPwdModeResponse struct{}"
	}

	return strings.Join([]string{"ResetPwdModeResponse", string(data)}, " ")
}
