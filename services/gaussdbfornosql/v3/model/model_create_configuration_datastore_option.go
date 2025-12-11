package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigurationDatastoreOption 数据库对象。
type CreateConfigurationDatastoreOption struct {

	// **参数解释：** 数据库类型。 **约束限制：** 不涉及。 **取值范围：** - GeminiDB Cassandra实例取值为“cassandra”。 - GeminiDB Mongo实例取值为\"mongodb\"。 - GeminiDB Influx实例取值为\"influxdb\"。 - GeminiDB Redis实例取值为\"redis\"。 - GeminiDB DynamoDB实例取值为\"dynamodb\"。 - GeminiDB HBase实例取值为\"hbase\"。 **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 数据库版本。 **约束限制：** 不涉及。 **取值范围：** - GeminiDB Cassandra实例支持3.11版本，取值为“3.11”。 - GeminiDB Mongo实例支持4.0版本，取值为\\\"4.0\\\"。 - GeminiDB Influx实例支持1.8版本，取值\\\"1.8\\\"。 - GeminiDB Redis实例支持5.0版本，取值\\\"5.0\\\"。 **默认取值：** 不涉及。
	Version string `json:"version"`

	// **参数解释：** 数据库实例类型。 **约束限制：** 当创建支持GeminiDB Mongo、GeminiDB Cassandra云原生部署模式实例的参数模板， 该参数必选。 **取值范围：** - GeminiDB Mongo副本集实例取值为\\\"ReplicaSet\\\"。 - GeminiDB Influx 单节点实例取值为\\\"InfluxdbSingle\\\"。 - GeminiDB Influx 集群增强版实例取值为\\\"EnhancedCluster\\\"。 - GeminiDB Cassandra云原生部署模式实例取值为\\\"CloudNativeCluster\\\"。 - GeminiDB Influx 云原生部署模式实例取值为\\\"CloudNativeCluster\\\"。 - GeminiDB Redis云原生部署模式实例取值为\\\"CloudNativeCluster\\\"。 **默认取值：** 不传该参数，对GeminiDB Cassandra实例默认创建支持经典部署模式实例的参数模板。
	Mode *string `json:"mode,omitempty"`
}

func (o CreateConfigurationDatastoreOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationDatastoreOption struct{}"
	}

	return strings.Join([]string{"CreateConfigurationDatastoreOption", string(data)}, " ")
}
