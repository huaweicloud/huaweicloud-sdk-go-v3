package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstanceRequestBody 创建云堡垒机实例变更任务请求体。  > 说明： new_resource_spec_code和attach_disk_size字段只能选择使用，不能同时使用。
type ChangeInstanceRequestBody struct {
	ServerId *interface{} `json:"server_id"`

	// 待变更云堡垒机规格ID，例如： - cbh.basic.50 - cbh.enhance.50  可参考接口\"查询云堡垒机规格信息\"获取
	NewResourceSpecCode *string `json:"new_resource_spec_code,omitempty"`

	// 附加磁盘大小。单位TB  > 说明： 附加磁盘和规格自带磁盘大小不能超过300TB。
	AttachDiskSize *int32 `json:"attach_disk_size,omitempty"`

	// 是否自动支付，下单订购后，是否自动从客户的华为云账户中支付，而不需要客户手动去进行支付。 - 1：是（会自动选择折扣和优惠券进行优惠，然后自动从客户华为云账户中支付），自动支付失败后会生成订单成功(该订单应付金额是优惠后金额)、但订单状态为“待支付”，等待客户手动支付(手动支付时，客户还可以修改系统自动选择的折扣和优惠券) - 0：否（需要客户手动去支付，客户可以选择折扣和优惠券。）  默认值为“0”
	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`
}

func (o ChangeInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeInstanceRequestBody", string(data)}, " ")
}
