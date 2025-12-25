package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Snapshots **参数解释**： 快照对象列表。 **取值范围**： 不涉及。
type Snapshots struct {

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

	// **参数解释**： 快照大小，单位 GB。 **取值范围**： 不涉及。
	Size float64 `json:"size"`

	// **参数解释**： 快照状态。 **取值范围**： - CREATING：创建中。 - AVAILABLE：可用。 - UNAVAILABLE：不可用。 - FROZEN： 普通冻结。 - POLICE_FROZEN： 公安冻结。
	Status string `json:"status"`

	// **参数解释**： 快照创建类型。 **取值范围**： - AUTO: 自动。 - MANUAL: 手动。
	Type string `json:"type"`

	// **参数解释**： 快照对应的集群ID。 **取值范围**： 不涉及。
	ClusterId string `json:"cluster_id"`

	Datastore *Datastore `json:"datastore,omitempty"`

	// **参数解释**： 集群名称。 **取值范围**： 不涉及。
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 更新时间。 **取值范围**： 不涉及。
	Updated *string `json:"updated,omitempty"`

	// **参数解释**： 快照期待开始执行时间。 **取值范围**： 不涉及。
	BakExpectedStartTime *string `json:"bak_expected_start_time,omitempty"`

	// **参数解释**： 保存天数。 **取值范围**： 不涉及。
	BakKeepDay *int32 `json:"bak_keep_day,omitempty"`

	// **参数解释**： 备份周期。 **取值范围**： 不涉及。
	BakPeriod *string `json:"bak_period,omitempty"`

	// **参数解释**： 数据库用户名。 **取值范围**： 不涉及。
	DbUser *string `json:"db_user,omitempty"`

	// **参数解释**： 进度。 **取值范围**： 不涉及。
	Progress *string `json:"progress,omitempty"`

	// **参数解释**： 备份产生的Key。 **取值范围**： 不涉及。
	BackupKey *SnapshotsBackupKey `json:"backup_key,omitempty"`

	// **参数解释**： 增量快照使用前一个快照的BackupKey。当取值为FULL时表示这是一个全量快照。 **取值范围**： 不涉及。
	PriorBackupKey *string `json:"prior_backup_key,omitempty"`

	// **参数解释**： 对应全量快照的BackupKey。 **取值范围**： 不涉及。
	BaseBackupKey *string `json:"base_backup_key,omitempty"`

	// **参数解释**： 备份介质。 **取值范围**： 不涉及。
	BackupDevice *string `json:"backup_device,omitempty"`

	// **参数解释**： 总备份文件大小。 **取值范围**： 不涉及。
	TotalBackupSize *int64 `json:"total_backup_size,omitempty"`

	// **参数解释**： 对应全量快照名称。 **取值范围**： 不涉及。
	BaseBackupName *string `json:"base_backup_name,omitempty"`

	// **参数解释**： 是否支持备份到当前集群。 **取值范围**： 不涉及。
	SupportInplaceRestore *bool `json:"support_inplace_restore,omitempty"`

	// **参数解释**： 是否支持细粒度备份。 **取值范围**： 不涉及。
	FineGrainedBackup *bool `json:"fine_grained_backup,omitempty"`

	// **参数解释**： 备份等级。 **取值范围**： 不涉及。
	BackupLevel *SnapshotsBackupLevel `json:"backup_level,omitempty"`

	FineGrainedBackupDetail *ExtFineGrainedSnapshotDetail `json:"fine_grained_backup_detail,omitempty"`

	// **参数解释**： agent版本。 **取值范围**： 不涉及。
	GuestAgentVersion *string `json:"guest_agent_version,omitempty"`

	// **参数解释**： 集群状态。 **取值范围**： 不涉及。
	ClusterStatus *string `json:"cluster_status,omitempty"`

	// **参数解释**： 集群任务状态。 **取值范围**： 不涉及。
	ClusterTaskStatus *string `json:"cluster_task_status,omitempty"`

	// **参数解释**： 是否支持细粒度跨版本恢复。 **取值范围**： 不涉及。
	SupportFineGrainedCrossVersionRestore *bool `json:"support_fine_grained_cross_version_restore,omitempty"`

	// **参数解释**： 是否支持细粒度异构恢复。 **取值范围**： 不涉及。
	SupportFineGrainedAsymmetricRestore *bool `json:"support_fine_grained_asymmetric_restore,omitempty"`
}

func (o Snapshots) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Snapshots struct{}"
	}

	return strings.Join([]string{"Snapshots", string(data)}, " ")
}

type SnapshotsBackupKey struct {
	value string
}

type SnapshotsBackupKeyEnum struct {
	E_20160823_035923 SnapshotsBackupKey
}

func GetSnapshotsBackupKeyEnum() SnapshotsBackupKeyEnum {
	return SnapshotsBackupKeyEnum{
		E_20160823_035923: SnapshotsBackupKey{
			value: "20160823_035923",
		},
	}
}

func (c SnapshotsBackupKey) Value() string {
	return c.value
}

func (c SnapshotsBackupKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SnapshotsBackupKey) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type SnapshotsBackupLevel struct {
	value string
}

type SnapshotsBackupLevelEnum struct {
	CLUSTER SnapshotsBackupLevel
	SCHEMA  SnapshotsBackupLevel
	TABLE   SnapshotsBackupLevel
}

func GetSnapshotsBackupLevelEnum() SnapshotsBackupLevelEnum {
	return SnapshotsBackupLevelEnum{
		CLUSTER: SnapshotsBackupLevel{
			value: "cluster",
		},
		SCHEMA: SnapshotsBackupLevel{
			value: "schema",
		},
		TABLE: SnapshotsBackupLevel{
			value: "table",
		},
	}
}

func (c SnapshotsBackupLevel) Value() string {
	return c.value
}

func (c SnapshotsBackupLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SnapshotsBackupLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
