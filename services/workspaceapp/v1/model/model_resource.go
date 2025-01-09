package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Resource 包周期资源对象。
type Resource struct {
	AddVolumes *AddVolumes `json:"add_volumes,omitempty"`

	// 支付后跳转的地址。
	CloudServiceConsoleUrl *string `json:"cloud_service_console_url,omitempty"`

	CreateServices *CreateServices `json:"create_services"`

	// 订购关系当前是否是自动续费：0 否 1 是。
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	// 订购周期数取值大于0
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 包周期订单订购周期类型：2：月；3：年；4：包小时（仅限带宽加油包购买场景使用）5：绝对时间；（追加附属资源场景使用，比如主机上追加云硬盘）6：一次性（chargingMode=1 一次性计费场景使用），必填。
	PeriodType *int32 `json:"period_type,omitempty"`

	// 订购数量。
	SubscriptionNum int32 `json:"subscription_num"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
