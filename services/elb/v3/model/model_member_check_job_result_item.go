package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberCheckJobResultItem 参数解释：分组检查项汇总。
type MemberCheckJobResultItem struct {

	// 参数解释：检查项名称
	Name *string `json:"name,omitempty"`

	// 参数解释：异常原因
	Reason *string `json:"reason,omitempty"`

	// 参数解释：重要级别，分为Major(严重)和Tips(提示)
	Severity *string `json:"severity,omitempty"`

	// 参数解释：检查类别，config表示配置检查
	Subject *string `json:"subject,omitempty"`

	// 参数解释：任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 参数解释：异常原因模板
	ReasonTemplate *string `json:"reason_template,omitempty"`

	// 参数解释：异常结果变量参数表，用于结合异常原因模板动态生成异常原因
	ReasonParams *[]string `json:"reason_params,omitempty"`
}

func (o MemberCheckJobResultItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberCheckJobResultItem struct{}"
	}

	return strings.Join([]string{"MemberCheckJobResultItem", string(data)}, " ")
}
