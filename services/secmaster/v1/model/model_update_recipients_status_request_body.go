package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecipientsStatusRequestBody 更新收件人状态请求体
type UpdateRecipientsStatusRequestBody struct {

	// 收件人邮箱
	EmailAddress string `json:"email_address"`
}

func (o UpdateRecipientsStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecipientsStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRecipientsStatusRequestBody", string(data)}, " ")
}
