package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStarRocksDataReplicationsRequest Request Object
type ListStarRocksDataReplicationsRequest struct {

	// StarRocks实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage string `json:"X-Language"`

	// 查询记录数。每页查询数据同步任务的数量。
	Limit *string `json:"limit,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *string `json:"offset,omitempty"`
}

func (o ListStarRocksDataReplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStarRocksDataReplicationsRequest struct{}"
	}

	return strings.Join([]string{"ListStarRocksDataReplicationsRequest", string(data)}, " ")
}
