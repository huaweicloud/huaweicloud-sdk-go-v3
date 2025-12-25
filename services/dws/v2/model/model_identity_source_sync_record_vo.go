package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdentitySourceSyncRecordVo **参数解释**： 同步记录详细信息。 **取值范围**： 不涉及。
type IdentitySourceSyncRecordVo struct {

	// **参数解释**： 记录ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 身份源ID。 **取值范围**： 不涉及。
	IdentitySourceId *string `json:"identity_source_id,omitempty"`

	// **参数解释**： 身份源类型。 **取值范围**： - ldap：目录服务数据源。 - oneaccess：聚合型数据源
	IdentitySourceType *string `json:"identity_source_type,omitempty"`

	// **参数解释**： 任务开始时间。 **取值范围**： 不涉及。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 任务结束时间。 **取值范围**： 不涉及。
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： 添加记录数。 **取值范围**： 大于等于0。
	AddCount *int32 `json:"add_count,omitempty"`

	// **参数解释**： 更新记录数。 **取值范围**： 大于等于0。
	UpdateCount *int32 `json:"update_count,omitempty"`

	// **参数解释**： 删除记录数。 **取值范围**： 大于等于0。
	DeleteCount *int32 `json:"delete_count,omitempty"`

	// **参数解释**： 失败记录数。 **取值范围**： 大于等于0。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// **参数解释**： 状态。 **取值范围**： - RUNNING：运行中。 - FINISHED：已完成。 - FAILED：失败。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 失败原因。 **取值范围**： 不涉及。
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o IdentitySourceSyncRecordVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdentitySourceSyncRecordVo struct{}"
	}

	return strings.Join([]string{"IdentitySourceSyncRecordVo", string(data)}, " ")
}
