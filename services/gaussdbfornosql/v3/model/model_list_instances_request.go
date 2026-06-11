package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesRequest Request Object
type ListInstancesRequest struct {

	// **参数解释：** 实例ID。 **约束限制：** 如果id以“*”起始，表示按照“*”后面的值模糊匹配，否则，按照实际填写的id精确匹配查询。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 实例名称。 **约束限制：** 如果name以“*”起始，表示按照“*”后面的值模糊匹配，否则，按照实际填写的name精确匹配查询。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 数据库类型。 **约束限制：** 不涉及。 **取值范围：**   - 取值为“cassandra”，表示查询GeminiDB Cassandra数据库实例。   - 取值为“mongodb”，表示查询GeminiDB Mongo数据库实例。   - 取值为“influxdb”，表示查询GeminiDB Influx数据库实例。   - 取值为“redis”，表示查询GeminiDB Redis数据库实例。   - 取值为“dynamodb”，表示查询GeminiDB兼容DynamoDB数据库实例。   - 取值为“hbase”，表示查询GeminiDB HBase数据库实例。 **默认取值：** 如果不传该参数，表示查询所有数据库实例。
	DatastoreType *string `json:"datastore_type,omitempty"`

	// **参数解释：** 实例类型。 **约束限制：** 不涉及。 **取值范围：**    -  取值为“Cluster”，表示GeminiDB Cassandra、GeminiDB Influx、GeminiDB Redis Proxy、GeminiDB HBase、GeminiDB兼容DynamoDB经典部署模式集群实例类型。     -  取值为“CloudNativeCluster”，表示GeminiDB Cassandra、GeminiDB Influx、GeminiDB Redis、GeminiDB HBase、GeminiDB兼容DynamoDB云原生部署模式集群实例类型。    -  取值为“RedisCluster”，表示GeminiDB Redis Cluster经典部署模式集群实例类型。    -  取值为“Replication”，表示GeminiDB Redis经典部署模式主备实例类型。     -  取值为“InfluxdbSingle”，表示GeminiDB Influx经典部署模式单节点实例类型。    -  取值为“EnhancedCluster”，表示GeminiDB Influx经典部署模式集群增强实例类型。    -  取值为“ReplicaSet”，表示GeminiDB Mongo副本集实例类型。    -  如果不传datastore_type参数，自动忽略该参数设置。 **默认取值：** 不涉及。
	Mode *string `json:"mode,omitempty"`

	// **参数解释：** 虚拟私有云ID。获取方法如下：    - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。    - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询VPC列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释：** 子网的网络ID。获取方法如下：    - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。    - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询子网列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释：** 索引位置偏移量 ，表示从指定project ID下最新的实例创建时间开始，按时间的先后顺序偏移offset条数据后查询对应的实例信息。 **约束限制：** 取值大于或等于0。 **取值范围：** 不涉及。 **默认取值：** 不传该参数时，查询偏移量默认为0，表示从最新的实例创建时间对应的实例开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 查询实例个数上限值。 **约束限制：** 不涉及。 **取值范围：** 1~100。 **默认取值：** 不传该参数时，默认查询前100条实例信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
