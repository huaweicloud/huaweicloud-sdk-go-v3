package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAuditLogResponseBodyAuditRecords struct {

	// 日志ID
	Id *string `json:"id,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群类型
	ClusterType *string `json:"cluster_type,omitempty"`

	// 操作
	Operation *string `json:"operation,omitempty"`

	// 时间
	Time *int64 `json:"time,omitempty"`

	// 操作状态
	OperateStatus *int32 `json:"operate_status,omitempty"`

	// 操作结果消息
	OperateMessage *string `json:"operate_message,omitempty"`

	// 审计状态
	AuditStatus *int32 `json:"audit_status,omitempty"`
}

func (o ShowAuditLogResponseBodyAuditRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditLogResponseBodyAuditRecords struct{}"
	}

	return strings.Join([]string{"ShowAuditLogResponseBodyAuditRecords", string(data)}, " ")
}
