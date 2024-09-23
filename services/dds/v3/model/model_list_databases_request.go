package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabasesRequest Request Object
type ListDatabasesRequest struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`

	// 索引位置偏移量。 取值大于或等于0。不传该参数时，查询偏移量默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询实例个数上限值。 取值范围：1~100。不传该参数时，默认查询前100条实例信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListDatabasesRequest", string(data)}, " ")
}
