package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValidateTokenReqDto struct {

	// Access Token。
	Token string `json:"token"`

	// 是否生成新的Token。 * true：生成新的token值 * false：不生成新的token值
	NeedGenNewToken bool `json:"needGenNewToken"`

	// 是否需要返回用户帐号信息（帐号、用户名称等信息）。
	NeedAccountInfo *bool `json:"needAccountInfo,omitempty"`
}

func (o ValidateTokenReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateTokenReqDto struct{}"
	}

	return strings.Join([]string{"ValidateTokenReqDto", string(data)}, " ")
}
