package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuditDumpRecord 审计日志
type AuditDumpRecord struct {

	// **参数解释**： 集群ID。 **取值范围**： 36位UUID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 执行时间。
	ExecutorTime *string `json:"executor_time,omitempty"`

	// 开始时间。
	BeginTime *string `json:"begin_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 桶名。
	BucketName *string `json:"bucket_name,omitempty"`

	// 前缀。
	LocationPrefix *string `json:"location_prefix,omitempty"`

	// 结果。
	Result *string `json:"result,omitempty"`

	// 失败原因。
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o AuditDumpRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditDumpRecord struct{}"
	}

	return strings.Join([]string{"AuditDumpRecord", string(data)}, " ")
}
