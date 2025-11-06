package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendUserEmailVerifyCodeResponse Response Object
type SendUserEmailVerifyCodeResponse struct {

	// **参数解释：** 发送邮箱验证码结果。 **取值范围：** 字符串长度不少于1，不超过1000。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SendUserEmailVerifyCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendUserEmailVerifyCodeResponse struct{}"
	}

	return strings.Join([]string{"SendUserEmailVerifyCodeResponse", string(data)}, " ")
}
