package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterUpdateRecordResp **参数解释**： 集群升级记录响应体。 **取值范围**： - update-cluster：升级集群。 - commit-cluster：提交集群。 - rollback-cluster：集群回滚。 - update-hotpatch：热补丁升级。 - rollback-hotpatch：热补丁回滚。
type ClusterUpdateRecordResp struct {

	// **参数解释**： 升级项目ID。 **取值范围**： 不涉及。
	ItemId *string `json:"item_id,omitempty"`

	// **参数解释**： 升级状态。 **取值范围**： - Waiting：待升级。 - Update_Running：升级中。 - Update_Success：升级成功。 - Update_Failure：升级失败。 - Rollback_Running：回滚中。 - Rollback_Failure：回滚失败。 - Commit_Running：提交中。 - Commit_Failure：提交失败。 - Completed：升级完成。 - Disable：不支持升级。 - Install_Running：热补丁安装中。 - Uninstall_Failure：热补丁卸载重。 - Commit_HotPatch_Failure：热补丁提交失败。
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
