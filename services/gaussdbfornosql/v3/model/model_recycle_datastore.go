package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecycleDatastore 数据库信息。
type RecycleDatastore struct {

	// 数据库类型。   - 取值为“cassandra”，表示GeminiDB Cassandra数据库实例。   - 取值为“mongodb”，表示GeminiDB Mongo数据库实例。   - 取值为“influxdb”，表示GeminiDB Influx数据库实例。   - 取值为“redis”，表示GeminiDB Redis数据库实例。
	Type string `json:"type"`

	// 数据库版本。
	Version string `json:"version"`
}

func (o RecycleDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleDatastore struct{}"
	}

	return strings.Join([]string{"RecycleDatastore", string(data)}, " ")
}
