package model

import (
	"encoding/json"

	"strings"
)

// 数据库对象。
type CreateConfigurationDatastoreOption struct {
	// 数据库类型。 - GaussDB(for Cassandra)实例取值为“cassandra”。 - GaussDB(for Mongo)实例取值为\"mongodb\"。 - GaussDB(for Influx)实例取值为\"influxdb\"。

	Type string `json:"type"`
	// 数据库版本。 - GaussDB(for Cassandra)实例支持3.11版本，取值为“3.11”。 - GaussDB(for Mongo)实例支持3.4、4.0版本，取值为\"3.4\"或\"4.0\"。 - GaussDB(for Influx)实例支持1.7版本，取值\"1.7\"。

	Version string `json:"version"`
	// 数据库部署模式，GaussDB(for Mongo)该参数必选。 - GaussDB(for Mongo) 集群实例取值为\"Sharding\"。 - GaussDB(for Mongo) 副本集实例取值为\"ReplicaSet\"。

	Mode *string `json:"mode,omitempty"`
}

func (o CreateConfigurationDatastoreOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateConfigurationDatastoreOption struct{}"
	}

	return strings.Join([]string{"CreateConfigurationDatastoreOption", string(data)}, " ")
}
