package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDisasterRecoveryDrillsRequest struct {
	// 保护组的ID。

	ServerGroupId *string `json:"server_group_id,omitempty"`
	// 容灾演练的名称。支持模糊查询。

	Name *string `json:"name,omitempty"`
	// 容灾演练的状态。

	Status *string `json:"status,omitempty"`
	// 演练虚拟私有云ID。

	DrillVpcId *string `json:"drill_vpc_id,omitempty"`
	// 每次请求返回结果个数限制，取值范围为[0,1000]的正整数，默认值为1000。

	Limit *int32 `json:"limit,omitempty"`
	// 每次请求开始的下标，即偏移量，默认值为0。offset必须为数字，不能为负数。

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDisasterRecoveryDrillsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDisasterRecoveryDrillsRequest struct{}"
	}

	return strings.Join([]string{"ListDisasterRecoveryDrillsRequest", string(data)}, " ")
}
