package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 回收备份实例信息。
type RecycleInstance struct {

	// 实例ID。
	Id *string `json:"id,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 实例类型。   - 取值为“Cluster”，表示GaussDB(for Cassandra)、GaussDB(for Influx)、GaussDB(for Redis)集群实例类型。   - 取值为“ReplicaSet”，表示GaussDB(for Mongo)副本集实例类型。
	Mode *string `json:"mode,omitempty"`

	Datastore *RecycleDatastore `json:"datastore,omitempty"`

	// 计费方式。 计费方式。   - prePaid：预付费，即包年/包月。   - postPaid：后付费，即按需付费。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 企业项目ID，取值为“0”，表示为default企业项目
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 备份ID。
	BackupId *string `json:"backup_id,omitempty"`

	// 实例创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 实例删除时间。
	DeletedAt *string `json:"deleted_at,omitempty"`

	// 回收备份保留截止时间。
	RetainedUntil *string `json:"retained_until,omitempty"`
}

func (o RecycleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleInstance struct{}"
	}

	return strings.Join([]string{"RecycleInstance", string(data)}, " ")
}
