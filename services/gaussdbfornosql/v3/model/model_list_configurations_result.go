package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationsResult 参数模板信息。
type ListConfigurationsResult struct {

	// 参数模板ID。
	Id string `json:"id"`

	// 参数模板名称。
	Name string `json:"name"`

	// 参数模板描述。
	Description *string `json:"description,omitempty"`

	// 数据库版本名称。
	DatastoreVersionName string `json:"datastore_version_name"`

	// 数据库名称。
	DatastoreName string `json:"datastore_name"`

	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800
	Created string `json:"created"`

	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Updated string `json:"updated"`

	// 数据库实例类型。 GeminiDB Cassandra集群类型为\"Cluster\"。 GeminiDB Mongo副本集类型为\"ReplicaSet\"。 GeminiDB Mongo集群类型为\"Sharding\"。 GeminiDB Influx集群类型为\"Cluster\"。 GeminiDB Influx单节点类型为\"InfluxdbSingle\"。 GeminiDB Redis集群类型为“Cluster”。 GeminiDB Redis主备类型为“Replication”。
	Mode string `json:"mode"`

	// 是否是用户自定义参数模板： - false，表示为系统默认参数模板。 - true，表示为用户自定义参数模板。
	UserDefined bool `json:"user_defined"`
}

func (o ListConfigurationsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsResult struct{}"
	}

	return strings.Join([]string{"ListConfigurationsResult", string(data)}, " ")
}
