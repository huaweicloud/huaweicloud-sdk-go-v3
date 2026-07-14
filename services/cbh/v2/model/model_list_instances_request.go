package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesRequest Request Object
type ListInstancesRequest struct {

	// 云堡垒机实例ID。（非必传，需要查询单个实例详情时传入）
	InstanceId *int64 `json:"instance_id,omitempty"`

	// 查询返回的记录数量上限，默认值 1000，即一次最多返回 1000 条实例记录。最小值为1，最大值为1000。
	Limit *string `json:"limit,omitempty"`

	// 查询的起始偏移量（索引位置），默认值 0，即从第 0 条记录开始查。最小值为0。
	Offset *string `json:"offset,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
