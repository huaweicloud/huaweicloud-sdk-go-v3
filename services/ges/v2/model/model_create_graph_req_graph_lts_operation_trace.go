package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGraphReqGraphLtsOperationTrace 图实例是否开启审计日志，默认不开启。
type CreateGraphReqGraphLtsOperationTrace struct {

	// 是否开启图审计，默认“false”。
	EnableAudit *bool `json:"enable_audit,omitempty"`

	// LTS日志组名称。
	AuditLogGroupName *string `json:"audit_log_group_name,omitempty"`
}

func (o CreateGraphReqGraphLtsOperationTrace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGraphReqGraphLtsOperationTrace struct{}"
	}

	return strings.Join([]string{"CreateGraphReqGraphLtsOperationTrace", string(data)}, " ")
}
