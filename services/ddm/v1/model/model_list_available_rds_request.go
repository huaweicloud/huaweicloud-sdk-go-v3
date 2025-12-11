package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableRdsRequest Request Object
type ListAvailableRdsRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0。取值必须为数字，且不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数：每页多少条。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAvailableRdsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableRdsRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableRdsRequest", string(data)}, " ")
}
