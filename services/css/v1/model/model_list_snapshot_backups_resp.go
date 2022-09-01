package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 快照对象。
type ListSnapshotBackupsResp struct {

	// 快照创建时间。
	Created *string `json:"created,omitempty" xml:"created"`

	Datastore *ListSnapshotBackupsDatastoreResp `json:"datastore,omitempty" xml:"datastore"`

	// 快照描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 快照ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 集群ID。
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId"`

	// 集群名字。
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName"`

	// 快照名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 快照状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 快照更新时间，格式为ISO8601：CCYY-MM-DDThh:mm:ss。
	Updated *string `json:"updated,omitempty" xml:"updated"`

	// 快照创建类型： - 0：表示自动创建。 - 1：表示手动创建。
	BackupType *string `json:"backupType,omitempty" xml:"backupType"`

	// 创建快照方式。
	BackupMethod *string `json:"backupMethod,omitempty" xml:"backupMethod"`

	// 快照开始执行时间。
	BackupExpectedStartTime *string `json:"backupExpectedStartTime,omitempty" xml:"backupExpectedStartTime"`

	// 快照保留时间。
	BackupKeepDay *int32 `json:"backupKeepDay,omitempty" xml:"backupKeepDay"`

	// 快照每天执行的时间点。
	BackupPeriod *string `json:"backupPeriod,omitempty" xml:"backupPeriod"`

	// 要备份的索引。
	Indices *string `json:"indices,omitempty" xml:"indices"`

	// 要备份的索引的总shard数。
	TotalShards *int32 `json:"totalShards,omitempty" xml:"totalShards"`

	// 备份失败的shard数。
	FailedShards *int32 `json:"failedShards,omitempty" xml:"failedShards"`

	// 快照的版本。
	Version *string `json:"version,omitempty" xml:"version"`

	// 快照恢复的状态。
	RestoreStatus *string `json:"restoreStatus,omitempty" xml:"restoreStatus"`

	// 快照开始执行的时间戳。
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime"`

	// 快照执行结束的时间戳。
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime"`

	// 保存快照数据的桶名。
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName"`
}

func (o ListSnapshotBackupsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotBackupsResp struct{}"
	}

	return strings.Join([]string{"ListSnapshotBackupsResp", string(data)}, " ")
}
