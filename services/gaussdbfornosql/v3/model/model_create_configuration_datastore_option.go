package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigurationDatastoreOption 数据库对象。
type CreateConfigurationDatastoreOption struct {

	// 数据库类型。 - GeminiDB Cassandra实例取值为“cassandra”。 - GeminiDB Mongo实例取值为\"mongodb\"。 - GeminiDB Influx实例取值为\"influxdb\"。 - GeminiDB Redis实例取值为\"redis\"。
	Type string `json:"type"`

	// 数据库版本。 - GeminiDB Cassandra实例支持3.11版本，取值为“3.11”。 - GeminiDB Mongo实例支持4.0版本，取值为\"4.0\"。 - GeminiDB Influx实例支持1.7版本，取值\"1.7\"。 - GeminiDB Redis实例支持5.0版本，取值\"5.0\"。
	Version string `json:"version"`

	// 数据库部署模式，GeminiDB Mongo该参数必选。 - GeminiDB Mongo 副本集实例取值为\"ReplicaSet\"。
	Mode *string `json:"mode,omitempty"`
}

func (o CreateConfigurationDatastoreOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationDatastoreOption struct{}"
	}

	return strings.Join([]string{"CreateConfigurationDatastoreOption", string(data)}, " ")
}
