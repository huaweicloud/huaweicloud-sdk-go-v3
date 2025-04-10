package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatastoreOption struct {

	// 数据库类型。 - 支持GeminiDB Cassandra，GeminiDB Mongo，GeminiDB Influx数据库实例。 - 取值为“cassandra”，表示创建GeminiDB Cassandra数据库实例。 - 取值为“mongodb”，表示创建GeminiDB Mongo数据库实例。 - 取值为“influxdb”，表示创建GeminiDB Influx数据库实例。 - 取值为“redis”，表示创建GeminiDB Redis数据库实例。 - 取值为“dynamodb”，表示创建GeminiDB DynamoDB数据库实例。 - 取值为“hbase”，表示创建GeminiDB HBase数据库实例。
	Type string `json:"type"`

	// 数据库版本。 - GeminiDB Cassandra实例支持3.11版本，取值为“3.11”。 - GeminiDB Mongo实例支持4.0版本，取值为4.0\"。 - GeminiDB Influ经典部署模式集群类型实例支持1.8版本，取值为“1.8”。 - GeminiDB Influx云原生部署模式集群类型实例支持1.7版本，取值为“1.7”。 - GeminiDB Redis实例支持5.0版本，取值为“5.0”。 - GeminiDB DynamoDB实例取值为“”。 - GeminiDB HBase实例取值为“”。
	Version string `json:"version"`

	// 存储引擎。 - GeminiDB Cassandra实例支持RocksDB存储引擎，取值为“rocksDB”。 - GeminiDB Mongo实例支持RocksDB存储引擎，取值为“rocksDB”。 - GeminiDB Influx实例支持RocksDB存储引擎，取值为“rocksDB”。 - GeminiDB Redis实例支持RocksDB存储引擎，取值为“rocksDB”。
	StorageEngine string `json:"storage_engine"`
}

func (o DatastoreOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreOption struct{}"
	}

	return strings.Join([]string{"DatastoreOption", string(data)}, " ")
}
