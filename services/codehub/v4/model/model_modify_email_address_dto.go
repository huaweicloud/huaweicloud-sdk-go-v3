package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyEmailAddressDto struct {

	// **参数解释：** 邮箱。 **取值范围：** 字符串长度不少于1，不超过1000。
	Email *string `json:"email,omitempty"`

	// **参数解释：** 验证码。 **取值范围：** 字符串长度不少于1，不超过1000。
	VerifyCode *string `json:"verify_code,omitempty"`
}

func (o ModifyEmailAddressDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyEmailAddressDto struct{}"
	}

	return strings.Join([]string{"ModifyEmailAddressDto", string(data)}, " ")
}
