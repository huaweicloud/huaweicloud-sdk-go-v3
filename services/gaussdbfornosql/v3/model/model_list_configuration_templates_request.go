package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationTemplatesRequest Request Object
type ListConfigurationTemplatesRequest struct {

	// 索引位置，偏移量。   - 从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）。   - 取值必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。 - 取值范围: 1~100。 - 不传该参数时，默认查询前100条信息。
	Limit *int32 `json:"limit,omitempty"`

	// 【参数解释】 数据库名称。 【约束限制】 不限制。 【取值范围】 cassandra：表示支持GeminiDB Cassandra实例。 redis：表示支持GeminiDB Redis实例。 influxdb：表示支持GeminiDB Influx实例。 mongodb： 表示支持GeminiDB Mongo实例。 【默认取值】 不传该参数，则表示查询所有数据库类型。
	DatastoreName *string `json:"datastore_name,omitempty"`

	// 【参数解释】 数据库实例类型。 【约束限制】 不限制。 【取值范围】 * 取值为“CloudNativeCluster”, 表示查询支持GeminiDB Cassandra云原生部署模式实例的参数模板。 * 取值为“Cluster”, 表示查询GeminiDB Cassandra经典部署模式实例、GeminiDB Influx经典部署模式实例、GeminiDB Redis Proxy集群经典部署模式实例支持的参数模板。 * 取值为“Replication”, 表示查询支持GeminiDB Redis经典部署模式主备类型实例的参数组。 * 取值为“InfluxdbSingle”, 表示查询支持GeminiDB Influx经典部署模式单节点类型实例的参数组。 * 取值为“ReplicaSet”, 表示查询支持GeminiDB Mongo副本集类型实例的参数组。 * 取值为“All”， 表示查询所有部署模式的参数模板。 【默认取值】 不传该参数，则表示查询经典部署模式实例支持的参数组。
	Mode *string `json:"mode,omitempty"`
}

func (o ListConfigurationTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationTemplatesRequest", string(data)}, " ")
}
