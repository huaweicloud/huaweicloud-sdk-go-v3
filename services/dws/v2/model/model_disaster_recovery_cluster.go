package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisasterRecoveryCluster **参数解释**： 容灾集群信息。 **取值范围**： 不涉及。
type DisasterRecoveryCluster struct {

	// **参数解释**： 容灾集群信息ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 容灾集群名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 容灾集群所在可用区。 **取值范围**： 不涉及。
	ClusterAz *string `json:"cluster_az,omitempty"`

	// **参数解释**： 容灾集群角色。 **取值范围**： 不涉及。
	Role *string `json:"role,omitempty"`

	// **参数解释**： 容灾集群所在region。 **取值范围**： 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**： 容灾集群状态。 **取值范围**： - backuping，备份运行中。 - restoring，恢复运行中。 - stopped，集群已停止。 - waiting，容灾运行中。 - abnormal，容灾异常。 - drDeleted，容灾已删除。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 容灾进度。 **取值范围**： 不涉及。
	Progress *string `json:"progress,omitempty"`

	// **参数解释**： 上一次容灾时间。 **取值范围**： 不涉及。
	LastSuccessTime *string `json:"last_success_time,omitempty"`

	// **参数解释**： OBS桶名称。 **取值范围**： 不涉及。
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`

	// **参数解释**： 数据库版本。 **取值范围**： 不涉及。
	DatastoreVersion *string `json:"datastore_version,omitempty"`

	// **参数解释**： 数据库类型。 **取值范围**： 不涉及。
	DatastoreType *string `json:"datastore_type,omitempty"`

	// **参数解释**： 磁盘容量。 **取值范围**： 不涉及。
	DiskCapacity *string `json:"disk_capacity,omitempty"`

	// **参数解释**： 磁盘使用率。 **取值范围**： 不涉及。
	DiskUsed *string `json:"disk_used,omitempty"`
}

func (o DisasterRecoveryCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRecoveryCluster struct{}"
	}

	return strings.Join([]string{"DisasterRecoveryCluster", string(data)}, " ")
}
