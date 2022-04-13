package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValidateTokenReqDto struct {
	// 登录用账号的token字符串

	Token string `json:"token"`
	// 是否生成新的token，内部使用参数。 true：生成新的token值。 false：不生成新的token值。

	NeedGenNewToken bool `json:"needGenNewToken"`
	// 是否需要返回用户可见帐号信息（帐号、用户姓名等信息）。

	NeedAccountInfo *bool `json:"needAccountInfo,omitempty"`
}

func (o ValidateTokenReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateTokenReqDto struct{}"
	}

	return strings.Join([]string{"ValidateTokenReqDto", string(data)}, " ")
}
