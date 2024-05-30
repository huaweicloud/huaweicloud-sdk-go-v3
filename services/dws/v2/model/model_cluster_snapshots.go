package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterSnapshots 集群快照对象列表
type ClusterSnapshots struct {

	// 快照ID。
	Id *string `json:"id,omitempty"`

	// 快照名称。
	Name *string `json:"name,omitempty"`

	// 快照描述。
	Description *string `json:"description,omitempty"`

	// 快照创建的日期时间，格式为 ISO8601: YYYY-MM-DDThh:mm:ssZ。
	Started *string `json:"started,omitempty"`

	// 快照大小，单位 GB。
	Size *float64 `json:"size,omitempty"`

	// 快照状态： - CREATING：创建中。 - AVAILABLE：可用。 - UNAVAILABLE：不可用。 - RESTORING：恢复中。 - FROZEN： 普通冻结。 - POLICE_FROZEN： 公安冻结。
	Status *string `json:"status,omitempty"`

	// 快照对应的集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	Datastore *Datastore `json:"datastore,omitempty"`

	// 快照对应的集群名称。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 快照更新时间。
	Updated *string `json:"updated,omitempty"`

	// 快照类型： - MANUAL :手动快照， - AUTO :自动快照。
	Type *string `json:"type,omitempty"`

	// 快照预计开始时间。
	BakExpectedStartTime *string `json:"bak_expected_start_time,omitempty"`

	// 快照保留天数。
	BakKeepDay *int32 `json:"bak_keep_day,omitempty"`

	// 快照策略。
	BakPeriod *string `json:"bak_period,omitempty"`

	// 数据库用户。
	DbUser *string `json:"db_user,omitempty"`

	// 快照进度。
	Progress *string `json:"progress,omitempty"`

	// 快照BakcupKey。
	BackupKey *string `json:"backup_key,omitempty"`

	// 增量快照，使用的前一个快照BakcupKey。
	PriorBackupKey *string `json:"prior_backup_key,omitempty"`

	// 对应全量快照BakcupKey。
	BaseBackupKey *string `json:"base_backup_key,omitempty"`

	// 备份介质。
	BackupDevice *string `json:"backup_device,omitempty"`

	// 累计快照大小。
	TotalBackupSize *int32 `json:"total_backup_size,omitempty"`

	// 对应全量快照名称。
	BaseBackupName *string `json:"base_backup_name,omitempty"`

	// 是否支持就地恢复。
	SupportInplaceRestore *bool `json:"support_inplace_restore,omitempty"`

	// 是否是细粒度备份。
	FineGrainedBackup *bool `json:"fine_grained_backup,omitempty"`

	// 备份级别。
	BackupLevel *string `json:"backup_level,omitempty"`

	FineGrainedBackupDetail *FineGrainedSnapshotDetail `json:"fine_grained_backup_detail,omitempty"`

	// guestAgent版本。
	GuestAgentVersion *string `json:"guest_agent_version,omitempty"`

	// 集群状态。
	ClusterStatus *string `json:"cluster_status,omitempty"`
}

func (o ClusterSnapshots) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterSnapshots struct{}"
	}

	return strings.Join([]string{"ClusterSnapshots", string(data)}, " ")
}
