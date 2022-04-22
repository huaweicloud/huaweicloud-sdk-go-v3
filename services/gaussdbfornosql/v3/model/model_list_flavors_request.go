package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListFlavorsRequest struct {

	// 实例所在区域。
	Region string `json:"region"`

	// 数据库类型。   - 取值为“cassandra”，表示查询GaussDB(for Cassandra)数据库实例支持的规格。   - 取值为“mongodb”，表示查询GaussDB(for Mongo)数据库实例支持的规格。   - 取值为“influxdb”，表示查询GaussDB(for Influx)数据库实例支持的规格。   - 如果不传该参数，默认为“cassandra”。
	EngineName *string `json:"engine_name,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}
