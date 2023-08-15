package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimSendReport 报表详情。
type AimSendReport struct {

	// 报表日期时间。
	ReportTime *string `json:"report_time,omitempty"`

	// 智能信息模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 提交号码总数量。
	SubmissionCount *int32 `json:"submission_count,omitempty"`

	// 支持解析数量。  > 此数据不包括通过API发送的智能信息任务。
	SupportResolveCount *int32 `json:"support_resolve_count,omitempty"`

	// 发送数量。
	SendCount *int32 `json:"send_count,omitempty"`

	// 成功发送数量。
	SendSuccessCount *int32 `json:"send_success_count,omitempty"`

	// 成功解析数量。
	ResolveSuccessCount *int32 `json:"resolve_success_count,omitempty"`
}

func (o AimSendReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimSendReport struct{}"
	}

	return strings.Join([]string{"AimSendReport", string(data)}, " ")
}
