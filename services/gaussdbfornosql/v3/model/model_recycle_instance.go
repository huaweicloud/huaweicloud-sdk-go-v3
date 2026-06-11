package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecycleInstance 回收备份实例信息。
type RecycleInstance struct {

	// **参数解释：** 实例ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 实例名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 实例类型。 **取值范围：** - 取值为“Cluster”，表示GeminiDB Cassandra、GeminiDB Influx、GeminiDB Redis Proxy经典部署模式集群实例类型。 - 取值为“CloudNativeCluster”，表示GeminiDB Cassandra、GeminiDB Influx、GeminiDB Redis云原生部署模式集群实例类型。 - 取值为“RedisCluster”，表示GeminiDB Redis Cluster经典部署模式集群实例类型。 - 取值为“Replication”，表示GeminiDB Redis经典部署模式主备实例类型。 - 取值为“InfluxdbSingle”，表示GeminiDB Influx经典部署模式单节点实例类型。 - 取值为“ReplicaSet”，表示GeminiDB Mongo副本集实例类型。
	Mode *string `json:"mode,omitempty"`

	// **参数解释：** 产品类型。 **取值范围：** GeminiDB Redis云原生部署模式集群涉及此字段，取值：   -  Standard 标准型   -  Capacity 容量型
	ProductType *string `json:"product_type,omitempty"`

	DataStore *RecycleDatastore `json:"data_store,omitempty"`

	// **参数解释：** 计费方式。 **取值范围：** - prePaid：预付费，即包年/包月。 - postPaid：后付费，即按需付费。
	ChargeType *string `json:"charge_type,omitempty"`

	// **参数解释：** 企业项目ID，取值为“0”，表示为default企业项目 **取值范围：** 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 备份ID。 **取值范围：** 不涉及。
	BackupId *string `json:"backup_id,omitempty"`

	// **参数解释：** 实例创建时间。 **取值范围：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 实例删除时间。 **取值范围：** 不涉及。
	DeletedAt *string `json:"deleted_at,omitempty"`

	// **参数解释：** 回收备份保留截止时间。 **取值范围：** 不涉及。
	RetainedUntil *string `json:"retained_until,omitempty"`
}

func (o RecycleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleInstance struct{}"
	}

	return strings.Join([]string{"RecycleInstance", string(data)}, " ")
}
