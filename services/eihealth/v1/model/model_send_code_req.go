package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendCodeReq 发送验证码请求体
type SendCodeReq struct {

	// 认证方式:sms，email，vmfa
	Method string `json:"method"`
}

func (o SendCodeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendCodeReq struct{}"
	}

	return strings.Join([]string{"SendCodeReq", string(data)}, " ")
}
