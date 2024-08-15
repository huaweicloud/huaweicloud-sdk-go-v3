package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthRef 参数解释：带宽对象ID。 约束限制： - 仅在创建或更新公网IPv6负载均衡器时有效。 - 若选择创建新EIP并指定共享带宽时，此EIP会被分进共享带宽里面。  [不支持IPv6，请勿使用。](tag:dt,dt_test)
type BandwidthRef struct {

	// 参数解释：共享带宽的ID。
	Id string `json:"id"`
}

func (o BandwidthRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthRef struct{}"
	}

	return strings.Join([]string{"BandwidthRef", string(data)}, " ")
}
