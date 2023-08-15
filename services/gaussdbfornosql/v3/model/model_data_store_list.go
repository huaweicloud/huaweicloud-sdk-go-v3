package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataStoreList struct {

	// 数据库引擎。
	DatastoreName string `json:"datastore_name"`

	// 数据库引擎版本。
	Version string `json:"version"`

	// 实例类型。 取值为“Cluster”，表示GaussDB(for Cassandra)、GaussDB(for Influx)、GaussDB(for Redis)集群实例类型。 取值为“InfluxdbSingle”，表示GaussDB(for Influx)单节点实例类型。 取值为“ReplicaSet”，表示GaussDB(for Mongo)副本集实例类型。
	Mode string `json:"mode"`
}

func (o DataStoreList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataStoreList struct{}"
	}

	return strings.Join([]string{"DataStoreList", string(data)}, " ")
}
