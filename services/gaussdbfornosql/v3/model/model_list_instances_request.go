package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListInstancesRequest struct {

	// 实例ID。 如果id以“*”起始，表示按照“*”后面的值模糊匹配，否则，按照实际填写的id精确匹配查询。
	Id *string `json:"id,omitempty"`

	// 实例名称。 如果name以“*”起始，表示按照“*”后面的值模糊匹配，否则，按照实际填写的name精确匹配查询。
	Name *string `json:"name,omitempty"`

	// 实例类型。   - 取值为“Cluster”，表示GaussDB(for Cassandra)集群实例类型。   - 取值为“Sharding”，表示GaussDB(for Mongo)集群实例类型。   - 取值为“ReplicaSet”，表示GaussDB(for Mongo)副本集实例类型。   - 取值为“InfluxdbCluster”，表示GaussDB(for Influx)集群实例类型。
	Mode *string `json:"mode,omitempty"`

	// 数据库类型。   - 取值为“cassandra”，表示查询GaussDB(for Cassandra)数据库实例。   - 取值为“mongodb”，表示查询GaussDB(for Mongo)数据库实例。   - 取值为“influxdb”，表示查询GaussDB(for Influx)数据库实例。   - 如果不传该参数，表示查询所有数据库实例。
	DatastoreType *string `json:"datastore_type,omitempty"`

	// 虚拟私有云ID，获取方法如下：   - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。   - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询VPC列表。
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网的网络ID，获取方法如下：   - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。   - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询子网列表。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 索引位置偏移量，表示从指定project ID下最新的实例创建时间开始，按时间的先后顺序偏移offset条数据后查询对应的实例信息。 取值大于或等于0。不传该参数时，查询偏移量默认为0，表示从最新的实例创建时间对应的实例开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 查询实例个数上限值。 取值范围：1~100。不传该参数时，默认查询前100条实例信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
