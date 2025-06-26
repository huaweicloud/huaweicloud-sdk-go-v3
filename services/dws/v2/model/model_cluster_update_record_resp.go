package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterUpdateRecordResp **参数解释**： 集群升级记录响应体。 **取值范围**： 不涉及。
type ClusterUpdateRecordResp struct {

	// **参数解释**： 升级项目ID。 **取值范围**： 不涉及。
	ItemId *string `json:"item_id,omitempty"`

	// **参数解释**： 升级状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 升级类型。 **取值范围**： 不涉及。
	RecordType *string `json:"record_type,omitempty"`

	// **参数解释**： 升级前版本。 **取值范围**： 不涉及。
	FromVersion *string `json:"from_version,omitempty"`

	// **参数解释**： 目标版本。 **取值范围**： 不涉及。
	ToVersion *string `json:"to_version,omitempty"`

	// **参数解释**： 开始时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 升级任务ID。 **取值范围**： 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 失败原因。 **取值范围**： 不涉及。
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o ClusterUpdateRecordResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterUpdateRecordResp struct{}"
	}

	return strings.Join([]string{"ClusterUpdateRecordResp", string(data)}, " ")
}
