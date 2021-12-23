package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetAuditlogPolicyRequestBody struct {
	// 审计日志保存天数，取值范围0，7~732。 - 取值0，表示关闭审计日志策略。 - 取值7~732，表示开启审计日志策略，并设置审计日志保存天数为该值。

	KeepDays int32 `json:"keep_days"`
	// 仅关闭审计日志策略时有效。 - true（默认），表示关闭审计日志策略的同时，保留历史审计日志。 - false，表示关闭审计日志策略的同时，删除已有的历史审计日志。

	ReserveAuditlogs *string `json:"reserve_auditlogs,omitempty"`
	// 仅打开审计日志策略时有效，并且为空时，默认全部。审计范围。请输入数据库或集合名称，多个库或集合请用英文逗号分隔。若名称中有英文逗号，请在逗号前添加“$”符号，用以区分分隔符。

	AuditScope *string `json:"audit_scope,omitempty"`
	// 仅打开审计日志策略时有效，并且为空时，默认全部。审计类型。支持insert，delete，update，query等。

	AuditTypes *[]string `json:"audit_types,omitempty"`
}

func (o SetAuditlogPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditlogPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetAuditlogPolicyRequestBody", string(data)}, " ")
}
