package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsOperationTraceReq struct {

	// 是否开启图审计，默认“false”。
	EnableAudit *bool `json:"enableAudit,omitempty" xml:"enableAudit"`

	// LTS日志组名称。
	AuditLogGroupName *string `json:"auditLogGroupName,omitempty" xml:"auditLogGroupName"`
}

func (o LtsOperationTraceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsOperationTraceReq struct{}"
	}

	return strings.Join([]string{"LtsOperationTraceReq", string(data)}, " ")
}
