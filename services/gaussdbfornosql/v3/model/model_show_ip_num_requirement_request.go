package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpNumRequirementRequest Request Object
type ShowIpNumRequirementRequest struct {

	// 创建实例或扩容节点的个数。最大支持输入200。
	NodeNum int32 `json:"node_num"`

	// 数据库引擎名称。没有传入实例ID的时候该字段为必传。 - 取值为“cassandra”，表示GeminiDB Cassandra数据库引擎。 - 取值为“mongodb”，表示GeminiDB Mongo数据库引擎。 - 取值为“influxdb”，表示GeminiDB Influx数据库引擎。 - 取值为“redis”，表示GeminiDB Redis数据库引擎。
	EngineName *string `json:"engine_name,omitempty"`

	// 实例类型。没有传入实例ID的时候该字段为必传。 - 取值为“Cluster”，表示GeminiDB Cassandra、GeminiDB Influx、GeminiDB Redis集群实例类型。 - 取值为“ReplicaSet”，表示GeminiDB Mongo副本集实例类型。
	InstanceMode *string `json:"instance_mode,omitempty"`

	// 实例Id，可以调用5.3.3 查询实例列表和详情接口获取。如果未申请实例，可以调用5.3.1 创建实例接口创建。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ShowIpNumRequirementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpNumRequirementRequest struct{}"
	}

	return strings.Join([]string{"ShowIpNumRequirementRequest", string(data)}, " ")
}
