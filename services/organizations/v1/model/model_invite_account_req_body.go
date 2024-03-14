package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InviteAccountReqBody InviteAccount 操作的请求体。
type InviteAccountReqBody struct {
	Target *TargetDto `json:"target"`

	// 给收件账号所有者的邮件中的附加信息。
	Notes string `json:"notes"`

	// 要绑定到新创建的账号的标签列表。
	Tags *[]TagDto `json:"tags,omitempty"`
}

func (o InviteAccountReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteAccountReqBody struct{}"
	}

	return strings.Join([]string{"InviteAccountReqBody", string(data)}, " ")
}
