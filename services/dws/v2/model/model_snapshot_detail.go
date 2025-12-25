package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SnapshotDetail **参数解释**： 快照详情对象。 **取值范围**： 不涉及。
type SnapshotDetail struct {

	// **参数解释**： 快照ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 快照名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 快照描述。 **取值范围**： 不涉及。
	Description string `json:"description"`

	// **参数解释**： 快照创建的日期时间，格式为 ISO8601: YYYY-MM-DDThh:mm:ssZ。 **取值范围**： 不涉及。
	Started string `json:"started"`

	// **参数解释**： 快照完成的日期时间，格式为 ISO8601: YYYY-MM-DDThh:mm:ssZ。 **取值范围**： 不涉及。
	Finished string `json:"finished"`

	// **参数解释**： 快照大小，单位GB。 **取值范围**： 不涉及。
	Size float64 `json:"size"`

	// **参数解释**： 快照状态： **取值范围**： - CREATING：创建中。 - AVAILABLE：可用。 - UNAVAILABLE：不可用。 - FROZEN：普通冻结。 - POLICE_FROZEN：公安冻结。
	Status string `json:"status"`

	// **参数解释**： 快照创建类型。 **取值范围**： - AUTO: 自动。 - MANUAL: 手动。
	Type string `json:"type"`

	// **参数解释**： 快照对应的集群ID。 **取值范围**： 不涉及。
	ClusterId string `json:"cluster_id"`

	Datastore *Datastore `json:"datastore,omitempty"`

	// **参数解释**： 快照对应的集群名称。 **取值范围**： 不涉及。
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 快照更新时间。 **取值范围**： 不涉及。
	Updated *string `json:"updated,omitempty"`

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

	// **参数解释**： 快照BakcupKey。 **取值范围**： 不涉及。
	BackupKey *string `json:"backup_key,omitempty"`

	// **参数解释**： 增量快照，使用的前一个快照BakcupKey。 **取值范围**： 不涉及。
	PriorBackupKey *string `json:"prior_backup_key,omitempty"`

	// **参数解释**： 对应全量快照BakcupKey。 **取值范围**： 不涉及。
	BaseBackupKey *string `json:"base_backup_key,omitempty"`

	// **参数解释**： 备份介质。 **取值范围**： NBU、OBS。
	BackupDevice *string `json:"backup_device,omitempty"`

	// **参数解释**： 累计快照大小。 **取值范围**： 不涉及。
	TotalBackupSize *int32 `json:"total_backup_size,omitempty"`

	// **参数解释**： 对应全量快照名称。 **取值范围**： 不涉及。
	BaseBackupName *string `json:"base_backup_name,omitempty"`

	// **参数解释**： 是否支持就地恢复。 **取值范围**： 不涉及。
	SupportInplaceRestore *bool `json:"support_inplace_restore,omitempty"`

	// **参数解释**： 是否是细粒度备份。 **取值范围**： 不涉及。
	FineGrainedBackup *bool `json:"fine_grained_backup,omitempty"`

	// **参数解释**： 备份级别。 **取值范围**： cluster：集群级快照； schema：schema级快照； table：表级快照；
	BackupLevel *string `json:"backup_level,omitempty"`

	FineGrainedBackupDetail *FineGrainedSnapshotDetail `json:"fine_grained_backup_detail,omitempty"`

	// **参数解释**： guestAgent版本。 **取值范围**： 不涉及。
	GuestAgentVersion *string `json:"guest_agent_version,omitempty"`

	// **参数解释**： 集群状态。 **取值范围**： - CREATING：创建中。 - AVAILABLE：可用。 - UNAVAILABLE：不可用。 - RESTORING：恢复中。 - FROZEN： 普通冻结。 - POLICE_FROZEN： 公安冻结。
	ClusterStatus *string `json:"cluster_status,omitempty"`

	// **参数解释**： 集群任务状态。 **取值范围**： 不涉及。
	ClusterTaskStatus *string `json:"cluster_task_status,omitempty"`

	// **参数解释**： 是否支持细粒度跨版本恢复。 **取值范围**： 不涉及。
	SupportFineGrainedCrossVersionRestore *bool `json:"support_fine_grained_cross_version_restore,omitempty"`

	// **参数解释**： 是否支持细粒度异构恢复。 **取值范围**： 不涉及。
	SupportFineGrainedAsymmetricRestore *bool `json:"support_fine_grained_asymmetric_restore,omitempty"`
}

func (o SnapshotDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotDetail struct{}"
	}

	return strings.Join([]string{"SnapshotDetail", string(data)}, " ")
}
