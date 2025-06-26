package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuditDumpRecord **参数解释**： 审计日志。 **取值范围**： 不涉及。
type AuditDumpRecord struct {

	// **参数解释**： 集群ID。 **取值范围**： 36位UUID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 执行时间。 **取值范围**： 不涉及。
	ExecutorTime *string `json:"executor_time,omitempty"`

	// **参数解释**： 开始时间。 **取值范围**： 不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**： 结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 桶名。 **取值范围**： 不涉及。
	BucketName *string `json:"bucket_name,omitempty"`

	// **参数解释**： 前缀。 **取值范围**： 不涉及。
	LocationPrefix *string `json:"location_prefix,omitempty"`

	// **参数解释**： 结果。 **取值范围**： 不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释**： 失败原因。 **取值范围**： 不涉及。
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o AuditDumpRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditDumpRecord struct{}"
	}

	return strings.Join([]string{"AuditDumpRecord", string(data)}, " ")
}
