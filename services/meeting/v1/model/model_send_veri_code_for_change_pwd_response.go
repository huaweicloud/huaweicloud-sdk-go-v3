package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SendVeriCodeForChangePwdResponse struct {

	// 过期时间，单位：秒。
	Expire *int32 `json:"expire,omitempty"`

	// 如果通过手机发送验证码，则该字段携带该用户绑定的手机号（手机号经过处理，屏蔽中间几位，如+8618****12345）。
	BindPhone *string `json:"bindPhone,omitempty"`

	// 如果通过邮箱发送验证码，则该字段携带用户绑定的邮箱帐号（邮箱帐号经过处理，屏蔽中间几位，如tes****ount@huawei.com）。
	BindEmail      *string `json:"bindEmail,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SendVeriCodeForChangePwdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendVeriCodeForChangePwdResponse struct{}"
	}

	return strings.Join([]string{"SendVeriCodeForChangePwdResponse", string(data)}, " ")
}
