package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataStoreList struct {

	// 数据库引擎。
	DatastoreName string `json:"datastore_name"`

	// 数据库引擎版本。
	Version string `json:"version"`

	// 实例类型。 取值为“Cluster”，表示GeminiDB Cassandra、GeminiDB Influx、GeminiDB Redis集群实例类型。 取值为“InfluxdbSingle”，表示GeminiDB Influx单节点实例类型。 取值为“ReplicaSet”，表示GeminiDB Mongo副本集实例类型。
	Mode string `json:"mode"`
}

func (o DataStoreList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataStoreList struct{}"
	}

	return strings.Join([]string{"DataStoreList", string(data)}, " ")
}
