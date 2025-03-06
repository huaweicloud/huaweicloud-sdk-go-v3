package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorInfosRequest Request Object
type ListFlavorInfosRequest struct {

	// 数据库类型。   - 取值为“cassandra”，表示查询GeminiDB Cassandra数据库实例支持的规格。   - 取值为“mongodb”，表示查询GeminiDB Mongo数据库实例支持的规格。   - 取值为“influxdb”，表示查询GeminiDB Influx数据库实例支持的规格。   - 取值为“redis”，表示查询GeminiDB Redis数据库实例支持的规格。   - 如果不传该参数，默认为“cassandra”。
	EngineName *string `json:"engine_name,omitempty"`

	// - 取值为“CloudNativeCluster”, 表示查询云原生部署模式支持的规格。 - 不传该参数表示查询数据库类型下的所有经典部署模式支持的规格。
	Mode *string `json:"mode,omitempty"`

	// 产品类型。   - Standard 标准型   - Capacity 容量型 当创建GeminiDB Redis云原生部署模式集群类型必传此参数。
	ProductType *string `json:"product_type,omitempty"`

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
