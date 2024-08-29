package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstanceTypeRequest Request Object
type ChangeInstanceTypeRequest struct {

	// 实例id
	ServerId string `json:"server_id"`

	// 可用分区名称。  可参考接口\"获取服务可用区\"获取
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 是否自动支付，下单订购后，是否自动从客户的华为云账户中支付，而不需要客户手动去进行支付。 - 1：是（会自动选择折扣和优惠券进行优惠，然后自动从客户华为云账户中支付），自动支付失败后会生成订单成功(该订单应付金额是优惠后金额)、但订单状态为“待支付”，等待客户手动支付(手动支付时，客户还可以修改系统自动选择的折扣和优惠券) - 0：否（需要客户手动去支付，客户可以选择折扣和优惠券。）  默认值为“0”
	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`
}

func (o ChangeInstanceTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceTypeRequest struct{}"
	}

	return strings.Join([]string{"ChangeInstanceTypeRequest", string(data)}, " ")
}
