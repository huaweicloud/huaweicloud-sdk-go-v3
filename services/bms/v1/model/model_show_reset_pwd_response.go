package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResetPwdResponse struct {
	// 是否支持重置密码。True：支持一键重置密码。False：不支持一键重置密码

	ResetpwdFlag   *string `json:"resetpwd_flag,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResetPwdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResetPwdResponse struct{}"
	}

	return strings.Join([]string{"ShowResetPwdResponse", string(data)}, " ")
}
