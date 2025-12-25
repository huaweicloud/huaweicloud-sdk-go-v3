package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateItemResp **参数解释**： 集群升级路径响应体。 **取值范围**： 不涉及。
type UpdateItemResp struct {

	// **参数解释**： 升级项ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 起始版本。 **取值范围**： 不涉及。
	From *string `json:"from,omitempty"`

	// **参数解释**： 目标版本。 **取值范围**： 不涉及。
	To *string `json:"to,omitempty"`

	// **参数解释**： 升级路径状态。 **取值范围**： - Waiting：待升级。 - Update_Running：升级中。 - Update_Success：升级成功。 - Update_Failure：升级失败。 - Rollback_Running：回滚中。 - Rollback_Failure：回滚失败。 - Commit_Running：提交中。 - Commit_Failure：提交失败。 - Completed：升级完成。 - Disable：不支持升级。 - Install_Running：热补丁安装中。 - Uninstall_Failure：热补丁卸载重。 - Commit_HotPatch_Failure：热补丁提交失败。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 升级进度。 **取值范围**： 不涉及。
	Process *string `json:"process,omitempty"`

	// **参数解释**： 起始时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 升级任务ID。 **取值范围**： 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 失败原因。 **取值范围**： 不涉及。
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o UpdateItemResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateItemResp struct{}"
	}

	return strings.Join([]string{"UpdateItemResp", string(data)}, " ")
}
