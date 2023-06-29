package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeVerifyReq 预验证请求体
type CodeVerifyReq struct {

	// 验证码
	Code string `json:"code"`

	// 认证方式
	Method string `json:"method"`
}

func (o CodeVerifyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeVerifyReq struct{}"
	}

	return strings.Join([]string{"CodeVerifyReq", string(data)}, " ")
}
