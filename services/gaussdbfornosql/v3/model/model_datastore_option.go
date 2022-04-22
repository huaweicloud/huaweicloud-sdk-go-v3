package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatastoreOption struct {

	// 数据库类型。 - 支持GaussDB(for Cassandra)，GaussDB(for Mongo)，GaussDB(for Influx)数据库实例。 - 取值为“cassandra”，表示创建GaussDB(for Cassandra)数据库实例。 - 取值为“mongodb”，表示创建GaussDB(for Mongo)数据库实例。 - 取值为“influxdb”，表示创建GaussDB(for Influx)数据库实例。
	Type string `json:"type"`

	// 数据库版本。 - GaussDB(for Cassandra)实例支持3.11版本，取值为“3.11”。 - GaussDB(for Mongo)实例支持3.4，4.0版本，取值为\"3.4\"或\"4.0\"。 - GaussDB(for Influx)实例支持1.7版本，取值为“1.7”。
	Version string `json:"version"`

	// 存储引擎。 - GaussDB(for Cassandra)实例支持RocksDB存储引擎，取值为“rocksDB”。 - GaussDB(for Mongo)实例支持RocksDB存储引擎，取值为“rocksDB”。 - GaussDB(for Influx)实例支持RocksDB存储引擎，取值为“rocksDB”。
	StorageEngine string `json:"storage_engine"`
}

func (o DatastoreOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreOption struct{}"
	}

	return strings.Join([]string{"DatastoreOption", string(data)}, " ")
}
