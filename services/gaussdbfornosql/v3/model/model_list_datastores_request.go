package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatastoresRequest Request Object
type ListDatastoresRequest struct {

	// 数据库类型。   - GeminiDB Cassandra数据库实例，取值为“cassandra”。   - GeminiDB Mongo数据库实例，取值为“mongodb”。   - GeminiDB Influx数据库实例，取值为“influxdb”。   - GeminiDB Redis数据库实例，取值为“redis”。
	DatastoreName string `json:"datastore_name"`
}

func (o ListDatastoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoresRequest struct{}"
	}

	return strings.Join([]string{"ListDatastoresRequest", string(data)}, " ")
}
