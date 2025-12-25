package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecipientsStatusRequestBody 查询收件人状态请求体
type ListRecipientsStatusRequestBody struct {

	// 收件人邮箱,可以为多个，多个时由英文分号(\";\")分割
	EmailAddress string `json:"email_address"`
}

func (o ListRecipientsStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecipientsStatusRequestBody struct{}"
	}

	return strings.Join([]string{"ListRecipientsStatusRequestBody", string(data)}, " ")
}
