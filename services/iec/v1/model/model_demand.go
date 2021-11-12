package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 租户需求
type Demand struct {
	// 所属运营商。

	Operator *string `json:"operator,omitempty"`
	// 站点需要发放的资源(组)总数。  > 实际发放实例数量为count*demand_count。

	DemandCount int32 `json:"demand_count"`
	// 线路ID。 多线路场景下，将在该线路下创建弹性公网IP。 > 覆盖规则为省级/大区时不支持指定线路ID创建边缘业务。

	PoolId *string `json:"pool_id,omitempty"`
}

func (o Demand) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Demand struct{}"
	}

	return strings.Join([]string{"Demand", string(data)}, " ")
}
