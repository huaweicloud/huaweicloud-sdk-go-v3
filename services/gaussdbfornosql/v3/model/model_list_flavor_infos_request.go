package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListFlavorInfosRequest struct {
	// 数据库类型。   - 取值为“cassandra”，表示查询GaussDB(for Cassandra)数据库实例支持的规格。   - 取值为“mongodb”，表示查询GaussDB(for Mongo)数据库实例支持的规格。   - 取值为“influxdb”，表示查询GaussDB(for Influx)数据库实例支持的规格。   - 取值为“redis”，表示查询GaussDB(for Redis)数据库实例支持的规格。   - 如果不传该参数，默认为“cassandra”。

	EngineName *string `json:"engine_name,omitempty"`
	// 索引位置，偏移量。   - 从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）。   - 取值必须为数字，不能为负数。

	Offset *int32 `json:"offset,omitempty"`
	// 查询规格信息上限值。   - 取值范围: 1~100。   - 不传该参数时，默认查询前100条规格信息。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFlavorInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorInfosRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorInfosRequest", string(data)}, " ")
}
