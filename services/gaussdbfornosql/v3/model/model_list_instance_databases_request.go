package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceDatabasesRequest Request Object
type ListInstanceDatabasesRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 索引位置，偏移量。    - 从第一条数据偏移offset条数据后开始查询，默认为0。   - 取值必须为数字，且不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。  - 取值范围：1~100。 - 不传该参数时，默认查询前100条信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceDatabasesRequest", string(data)}, " ")
}
