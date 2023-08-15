package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsOperationTraceReq struct {

	// 是否开启图审计，默认“false”。
	EnableAudit *bool `json:"enableAudit,omitempty"`

	// LTS日志组名称。
	AuditLogGroupName *string `json:"auditLogGroupName,omitempty"`
}

func (o LtsOperationTraceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsOperationTraceReq struct{}"
	}

	return strings.Join([]string{"LtsOperationTraceReq", string(data)}, " ")
}
