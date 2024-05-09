package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStarRocksDataBasesRequest Request Object
type ListStarRocksDataBasesRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage string `json:"X-Language"`

	// StarRocks实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 查询记录数
	Limit *string `json:"limit,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *string `json:"offset,omitempty"`

	// 查询的数据库名称，支持模糊搜索。
	DatabaseName *string `json:"database_name,omitempty"`
}

func (o ListStarRocksDataBasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStarRocksDataBasesRequest struct{}"
	}

	return strings.Join([]string{"ListStarRocksDataBasesRequest", string(data)}, " ")
}
