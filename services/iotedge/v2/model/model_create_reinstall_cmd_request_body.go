package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateReinstallCmdRequestBody struct {

	// 边缘节点设备密钥，如果不输入则平台随机生成
	DeviceSecret *string `json:"device_secret,omitempty"`

	// 边缘节点注册使用的验证码，如果不输入则平台随机生成。
	VerifyCode *string `json:"verify_code,omitempty"`
}

func (o CreateReinstallCmdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReinstallCmdRequestBody struct{}"
	}

	return strings.Join([]string{"CreateReinstallCmdRequestBody", string(data)}, " ")
}
