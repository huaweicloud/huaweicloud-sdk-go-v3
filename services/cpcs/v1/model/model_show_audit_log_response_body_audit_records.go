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
	Status *string `json:"status,omitempty"`

	// 操作失败消息
	FailureMessage *string `json:"failure_message,omitempty"`

	// 操作验证信息
	Verification *string `json:"verification,omitempty"`
}

func (o ShowAuditLogResponseBodyAuditRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditLogResponseBodyAuditRecords struct{}"
	}

	return strings.Join([]string{"ShowAuditLogResponseBodyAuditRecords", string(data)}, " ")
}
