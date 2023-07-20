package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotasRequest Request Object
type ShowQuotasRequest struct {

	// 数据库类型。 取值为“cassandra”，表示查询GaussDB(for Cassandra)数据库实例配额。 取值为“mongodb”，表示GaussDB(for Mongo)查询数据库实例配额。 取值为“influxdb”，表示查询GaussDB(for Influx)数据库实例配额。 取值为“redis”，表示查询GaussDB(for Redis)数据库实例配额。 如果不传该参数，表示查询所有数据库实例配额。
	DatastoreType *string `json:"datastore_type,omitempty"`

	// 实例类型。 取值为“Cluster”，表示GaussDB(for Cassandra)、GaussDB(for Influx)、GaussDB(for Redis)集群实例类型。 取值为“InfluxdbSingle”，表示GaussDB(for Influx)单节点实例类型。 取值为“ReplicaSet”，表示GaussDB(for Mongo)副本集实例类型。 如果不传datastore_type参数，自动忽略该参数设置，传入datastore_type时，该参数必填。
	Mode *string `json:"mode,omitempty"`
}

func (o ShowQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotasRequest", string(data)}, " ")
}
