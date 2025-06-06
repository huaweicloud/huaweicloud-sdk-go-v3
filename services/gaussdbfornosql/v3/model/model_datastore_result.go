package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatastoreResult 数据库信息，与请求参数相同。
type DatastoreResult struct {

	// 数据库类型。 - 支持GeminiDB Cassandra，GeminiDB Mongo，GeminiDB Influx数据库实例。 - 取值为“cassandra”，表示创建GeminiDB Cassandra数据库实例。 - 取值为“mongodb”，表示创建GeminiDB Mongo数据库实例。 - 取值为“influxdb”，表示创建GeminiDB Influx数据库实例。 - 取值为“dynamodb”，表示创建GeminiDB DynamoDB数据库实例。 - 取值为“hbase”，表示创建GeminiDB Hbase数据库实例。
	Type *string `json:"type,omitempty"`

	// 数据库版本。 - GeminiDB Cassandra实例支持3.11版本，取值为“3.11”。 - GeminiDB Mongo实例支持3.4，4.0版本，取值为\"3.4\"或\"4.0\"。 - GeminiDB Influx实例支持1.7版本，取值为“1.7”。 - GeminiDB DynamoDB取值为“”。 - GeminiDB HBase取值为“”。
	Version *string `json:"version,omitempty"`

	// 存储引擎。 - GeminiDB Cassandra实例支持RocksDB存储引擎，取值为“rocksDB”。 - GeminiDB Mongo实例支持RocksDB存储引擎，取值为“rocksDB”。 - GeminiDB Influx实例支持RocksDB存储引擎，取值为“rocksDB”。
	StorageEngine *string `json:"storage_engine,omitempty"`
}

func (o DatastoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreResult struct{}"
	}

	return strings.Join([]string{"DatastoreResult", string(data)}, " ")
}
