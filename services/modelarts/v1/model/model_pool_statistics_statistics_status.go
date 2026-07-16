package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolStatisticsStatisticsStatus **参数描述**： 不同状态下的资源池统计信息。
type PoolStatisticsStatisticsStatus struct {

	// **参数描述**： 正在创建中的资源池统计信息。 **取值范围**： 不涉及。
	Creating *int32 `json:"creating,omitempty"`

	// **参数描述**： 创建成功的资源池数量。 **取值范围**： 不涉及。
	Created *int32 `json:"created,omitempty"`

	// **参数描述**： 最近三天内创建失败的资源池数量，最大值为500。 **取值范围**： 不涉及。
	Failed *int32 `json:"failed,omitempty"`

	// **参数描述**： 等待中的资源池数量，通常是未支付的包周期资源池。 **取值范围**： 不涉及。
	Pending *int32 `json:"pending,omitempty"`
}

func (o PoolStatisticsStatisticsStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolStatisticsStatisticsStatus struct{}"
	}

	return strings.Join([]string{"PoolStatisticsStatisticsStatus", string(data)}, " ")
}
