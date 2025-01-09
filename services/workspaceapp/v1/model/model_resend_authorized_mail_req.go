package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResendAuthorizedMailReq 重发授权邮件的请求结构体（根据授权记录）。
type ResendAuthorizedMailReq struct {

	// 授权记录ID列表。
	Records []string `json:"records"`
}

func (o ResendAuthorizedMailReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResendAuthorizedMailReq struct{}"
	}

	return strings.Join([]string{"ResendAuthorizedMailReq", string(data)}, " ")
}
