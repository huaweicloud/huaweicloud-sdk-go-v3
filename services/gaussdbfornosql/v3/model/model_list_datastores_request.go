package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDatastoresRequest struct {
	// 数据库类型。 - GaussDB(for Cassandra)数据库实例，取值为“cassandra”。 - GaussDB(for Mongo)数据库实例，取值为“mongodb”。 - GaussDB(for Influx)数据库实例，取值为“influxdb”。

	DatastoreName string `json:"datastore_name"`
}

func (o ListDatastoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoresRequest struct{}"
	}

	return strings.Join([]string{"ListDatastoresRequest", string(data)}, " ")
}
