package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDatastoresRequest struct {

	// 数据库类型。   - GaussDB(for Cassandra)数据库实例，取值为“cassandra”。   - GaussDB(for Mongo)数据库实例，取值为“mongodb”。   - GaussDB(for Influx)数据库实例，取值为“influxdb”。   - GaussDB(for Redis)数据库实例，取值为“redis”。
	DatastoreName string `json:"datastore_name" xml:"datastore_name"`
}

func (o ListDatastoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoresRequest struct{}"
	}

	return strings.Join([]string{"ListDatastoresRequest", string(data)}, " ")
}
