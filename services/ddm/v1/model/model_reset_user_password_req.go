package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetUserPasswordReq ResetUserPasswordReq。
type ResetUserPasswordReq struct {

	// 重置后的新密码。
	Password string `json:"password"`
}

func (o ResetUserPasswordReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetUserPasswordReq struct{}"
	}

	return strings.Join([]string{"ResetUserPasswordReq", string(data)}, " ")
}
