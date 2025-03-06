package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecycleInstance 回收备份实例信息。
type RecycleInstance struct {

	// 实例ID。
	Id *string `json:"id,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 实例类型。   - 取值为“Cluster”，表示GeminiDB Cassandra、GeminiDB Influx、GeminiDB Redis 经典部署模式Proxy 集群实例类型。   - 取值为“CloudNativeCluster”，表示GeminiDB Cassandra、GeminiDB Influx、GeminiDB Redis云原生部署模式集群实例类型。   - 取值为“RedisCluster”，表示GeminiDB Redis经典部署模式Cluster集群实例类型。   - 取值为“Replication”，表示GeminiDB Redis经典部署模式主备实例类型。   - 取值为“InfluxdbSingle”，表示GeminiDB Influx 经典部署模式单节点实例类型。   - 取值为“ReplicaSet”，表示GeminiDB Mongo副本集实例类型。
	Mode *string `json:"mode,omitempty"`

	// 产品类型。 GeminiDB Redis云原生部署模式集群涉及此字段，取值：   -  Standard 标准型   -  Capacity 容量型
	ProductType *string `json:"product_type,omitempty"`

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
