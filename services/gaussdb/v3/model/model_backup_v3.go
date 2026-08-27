package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupV3 实例备份详细信息
type BackupV3 struct {

	// **参数解释**： 备份ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 备份描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 实例名称。 **取值范围**： 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**： 备份名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 备份大小。 **取值范围**： 不涉及。
	Size *string `json:"size,omitempty"`

	// **参数解释**： 大小单位（KB）。 **取值范围**： 不涉及。
	SizeUnit *string `json:"size_unit,omitempty"`

	// **参数解释**： 备份状态。 **取值范围**： - BUILDING：备份中。 - COMPLETED：备份完成。 - FAILED：备份失败。 - DELETING：备份删除中。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	Created *string `json:"created,omitempty"`

	// **参数解释**： 更新时间。 **取值范围**： 不涉及。
	Updated *string `json:"updated,omitempty"`

	// **参数解释**： 备份类型。 **取值范围**： - Db：物理备份。 - Snapshot：快照备份。
	BackupType *string `json:"backup_type,omitempty"`

	// **参数解释**： 备份级别。 **取值范围**： - 1：一级备份。 - 2：二级备份。
	BackupLevel *string `json:"backup_level,omitempty"`

	// **参数解释**： 备份方法。 **取值范围**： - Db：物理备份。 - Snapshot：快照备份。
	BackupMethod *string `json:"backup_method,omitempty"`

	// **参数解释**： 使用详情。 **取值范围**： 不涉及。
	UseDetail *string `json:"use_detail,omitempty"`

	// **参数解释**： UTC时区。 **取值范围**： 不涉及。
	TimeZone *string `json:"time_zone,omitempty"`

	// **参数解释**： 备份类型。 **取值范围**： - differential：差量备份。 - completed：全量备份。
	BackupMode *string `json:"backup_mode,omitempty"`
}

func (o BackupV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupV3 struct{}"
	}

	return strings.Join([]string{"BackupV3", string(data)}, " ")
}
