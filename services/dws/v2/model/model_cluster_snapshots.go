package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterSnapshots **参数解释**： 集群快照对象列表。 **取值范围**： 不涉及。
type ClusterSnapshots struct {

	// **参数解释**： 快照ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 快照名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 快照描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 快照创建的日期时间，格式为 ISO8601: YYYY-MM-DDThh:mm:ssZ。 **取值范围**： 不涉及。
	Started *string `json:"started,omitempty"`

	// **参数解释**： 快照大小，单位 GB。 **取值范围**： 不涉及。
	Size *float64 `json:"size,omitempty"`

	// **参数解释**： 快照状态。 **取值范围**： CREATING：创建中。 AVAILABLE：可用。 UNAVAILABLE：不可用。 RESTORING：恢复中。 FROZEN： 普通冻结。 POLICE_FROZEN： 公安冻结。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 快照对应的集群ID。 **取值范围**： 不涉及。
	ClusterId *string `json:"cluster_id,omitempty"`

	Datastore *Datastore `json:"datastore,omitempty"`

	// **参数解释**： 快照对应的集群名称。 **取值范围**： 不涉及。
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 快照更新时间。 **取值范围**： 不涉及。
	Updated *string `json:"updated,omitempty"`

	// **参数解释**： 快照类型。 **取值范围**： MANUAL：手动快照。 AUTO：自动快照。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 快照预计开始时间。 **取值范围**： 不涉及。
	BakExpectedStartTime *string `json:"bak_expected_start_time,omitempty"`

	// **参数解释**： 快照保留天数。 **取值范围**： 不涉及。
	BakKeepDay *int32 `json:"bak_keep_day,omitempty"`

	// **参数解释**： 快照策略。 **取值范围**： 不涉及。
	BakPeriod *string `json:"bak_period,omitempty"`

	// **参数解释**： 数据库用户。 **取值范围**： 不涉及。
	DbUser *string `json:"db_user,omitempty"`

	// **参数解释**： 快照进度。 **取值范围**： 不涉及。
	Progress *string `json:"progress,omitempty"`

	// **参数解释**： 快照备份key。 **取值范围**： 不涉及。
	BackupKey *string `json:"backup_key,omitempty"`

	// **参数解释**： 增量快照，使用的前一个快照备份key。 **取值范围**： 不涉及。
	PriorBackupKey *string `json:"prior_backup_key,omitempty"`

	// **参数解释**： 对应全量快照备份key。 **取值范围**： 不涉及。
	BaseBackupKey *string `json:"base_backup_key,omitempty"`

	// **参数解释**： 备份介质。 **取值范围**： 不涉及。
	BackupDevice *string `json:"backup_device,omitempty"`

	// **参数解释**： 累计快照大小。 **取值范围**： 不涉及。
	TotalBackupSize *int32 `json:"total_backup_size,omitempty"`

	// **参数解释**： 对应全量快照名称。 **取值范围**： 不涉及。
	BaseBackupName *string `json:"base_backup_name,omitempty"`

	// **参数解释**： 是否支持就地恢复。 **取值范围**： 不涉及。
	SupportInplaceRestore *bool `json:"support_inplace_restore,omitempty"`

	// **参数解释**： 是否是细粒度备份。 **取值范围**： 不涉及。
	FineGrainedBackup *bool `json:"fine_grained_backup,omitempty"`

	// **参数解释**： 备份级别。 **取值范围**： 不涉及。
	BackupLevel *string `json:"backup_level,omitempty"`

	FineGrainedBackupDetail *FineGrainedSnapshotDetail `json:"fine_grained_backup_detail,omitempty"`

	// **参数解释**： guestAgent版本。 **取值范围**： 不涉及。
	GuestAgentVersion *string `json:"guest_agent_version,omitempty"`

	// **参数解释**： 集群状态。 **取值范围**： 不涉及。
	ClusterStatus *string `json:"cluster_status,omitempty"`
}

func (o ClusterSnapshots) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterSnapshots struct{}"
	}

	return strings.Join([]string{"ClusterSnapshots", string(data)}, " ")
}
