package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WeakPasswordReq 弱密码请求参数。
type WeakPasswordReq struct {

	// 待测试是否是弱密码的字符串。
	Password string `json:"password"`
}

func (o WeakPasswordReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeakPasswordReq struct{}"
	}

	return strings.Join([]string{"WeakPasswordReq", string(data)}, " ")
}
