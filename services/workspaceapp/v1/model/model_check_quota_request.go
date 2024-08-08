package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckQuotaRequest Request Object
type CheckQuotaRequest struct {

	// 产品id。
	ProductId string `json:"product_id"`

	// 规格id。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 订单需要创建总实例数、订购数量。
	SubscriptionNum int32 `json:"subscription_num"`

	// 单台实例所需的磁盘大小（最大系统盘1块*1024、数据盘10块*32768）。
	DiskSize int32 `json:"disk_size"`

	// 单台实例所需的磁盘数量（最大系统盘1块、数据盘10块）。
	DiskNum int32 `json:"disk_num"`

	// 是否包周期。
	IsPeriod *bool `json:"is_period,omitempty"`

	// 主机id。
	DehId *string `json:"deh_id,omitempty"`

	// 云专属分布式存储池id。
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o CheckQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckQuotaRequest struct{}"
	}

	return strings.Join([]string{"CheckQuotaRequest", string(data)}, " ")
}
