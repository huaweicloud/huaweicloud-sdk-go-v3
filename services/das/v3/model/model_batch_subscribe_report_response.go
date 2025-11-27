package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSubscribeReportResponse Response Object
type BatchSubscribeReportResponse struct {

	// 成功的邮件模板列表
	SuccessEmailTemplateIds *[]int32 `json:"success_email_template_ids,omitempty"`

	// 失败的邮件模板列表
	FailedEmailTemplateIds *[]int32 `json:"failed_email_template_ids,omitempty"`
	HttpStatusCode         int      `json:"-"`
}

func (o BatchSubscribeReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSubscribeReportResponse struct{}"
	}

	return strings.Join([]string{"BatchSubscribeReportResponse", string(data)}, " ")
}
