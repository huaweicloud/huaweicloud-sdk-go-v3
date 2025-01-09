package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResendAuthorizationMailReq 重发授权邮件的请求结构体（根据授权邮件记录）。
type ResendAuthorizationMailReq struct {

	// 邮件记录。
	Records []AuthorizationMail `json:"records"`
}

func (o ResendAuthorizationMailReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResendAuthorizationMailReq struct{}"
	}

	return strings.Join([]string{"ResendAuthorizationMailReq", string(data)}, " ")
}
