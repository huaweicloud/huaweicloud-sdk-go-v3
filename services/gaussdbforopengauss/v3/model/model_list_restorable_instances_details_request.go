package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestorableInstancesDetailsRequest Request Object
type ListRestorableInstancesDetailsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *string `json:"X-Language,omitempty"`

	// 源实例id，需要恢复的实例ID。
	SourceInstanceId string `json:"source_instance_id"`

	// 实例备份信息ID，根据备份ID查询实例拓扑信息，过滤查询出来的实例，包含节点数，副本数等。参数为空时，根据restore_time查询。。
	BackupId *string `json:"backup_id,omitempty"`

	// 恢复点，当备份ID为空时，通过此参数查询实例拓扑信息，过滤实例列表。
	RestoreTime *string `json:"restore_time,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 查备份恢复的粒度。 **约束限制**: 不涉及。 **取值范围**:   - INSTANCE   - DATABASE_TABLE   - DATABASE **默认取值**: INSTANCE
	BackupRestoreType *string `json:"backup_restore_type,omitempty"`

	// **参数解释**: 源实例的备份类型。 **约束限制**: 不涉及。 **取值范围**:   - INSTANCE   - DATABASE_TABLE **默认取值**: INSTANCE
	SourceBackupSchema *string `json:"source_backup_schema,omitempty"`

	// **参数解释**: 目标实例ID，通过此参数过滤实例列表。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	TargetInstanceId *string `json:"target_instance_id,omitempty"`

	// **参数解释**: 目标实例名称，通过此参数过滤实例列表。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`
}

func (o ListRestorableInstancesDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestorableInstancesDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListRestorableInstancesDetailsRequest", string(data)}, " ")
}
