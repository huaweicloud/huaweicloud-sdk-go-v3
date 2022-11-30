package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库信息。
type RecycleDatastore struct {

	// 数据库类型。   - 取值为“cassandra”，表示GaussDB(for Cassandra)数据库实例。   - 取值为“mongodb”，表示GaussDB(for Mongo)数据库实例。   - 取值为“influxdb”，表示GaussDB(for Influx)数据库实例。   - 取值为“redis”，表示GaussDB(for Redis)数据库实例。
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
