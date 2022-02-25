package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlideVerifyCodeSendDto struct {
	// 用户身份信息（手机号码或邮箱账号或用户真实账号） maxLength：255 minLength：1

	User string `json:"user"`
	// 登录客户端类型。 * 0：Web客户端类型； * 5：cloudlink pc； * 6：cloudlink mobile； * 16：workplace pc

	ClientType int32 `json:"clientType"`
	// 校验类型。 * 0：登录； * 1：忘记密码； 默认值：0

	CheckType *int32 `json:"checkType,omitempty"`
}

func (o SlideVerifyCodeSendDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlideVerifyCodeSendDto struct{}"
	}

	return strings.Join([]string{"SlideVerifyCodeSendDto", string(data)}, " ")
}
