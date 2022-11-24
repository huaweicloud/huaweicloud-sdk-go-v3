package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdminResetPwdReqDto struct {

	// 被修改密码的用户帐号。
	Account string `json:"account"`
}

func (o AdminResetPwdReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdminResetPwdReqDto struct{}"
	}

	return strings.Join([]string{"AdminResetPwdReqDto", string(data)}, " ")
}
