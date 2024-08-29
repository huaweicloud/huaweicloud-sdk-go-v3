package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceBody 创建堡垒机实例请求参数。
type CreateInstanceBody struct {

	// 待创建云堡垒机规格ID，例如： - cbh.basic.50 - cbh.enhance.50  可参考接口\"查询云堡垒机规格信息\"获取。
	Specification string `json:"specification"`

	// 云堡垒机实例名称，取值范围：  只能由中文字符、英文字母、数字及“_”、“-”组成，且长度为[1-64]个字符。  例如：CBH-6b8e
	InstanceName string `json:"instance_name"`

	// 堡垒机实例前端登录密码。  密码规则：8-32位、不能包含amdin或nidma及其大写形式、必须包含大小写数字特殊字符（!@$%^-_=+[{}]:,./?~#*）四种类型中的三种、不能包含超过2个连续的相同字符（区分大小写）。
	Password string `json:"password"`

	// 创建云堡垒机实例所在局点ID。   可参考接口\"查询云堡垒机规格信息\"获取
	Region string `json:"region"`

	// 创建云堡垒机所在的可用分区，需要指定可用分区名称。(主备模式是作为主机可用区)  可参考接口\"获取服务可用区\"获取
	AvailabilityZone string `json:"availability_zone"`

	// 创建云堡垒机备机所在的可用分区，需要指定可用分区名称。(只创建单机时不传此字段)。  可参考接口\"获取服务可用区\"获取
	SlaveAvailabilityZone *string `json:"slave_availability_zone,omitempty"`

	// 计费模式。 - 0 包周期计费。
	ChargingMode int32 `json:"charging_mode"`

	// 订购周期类型。（包周期模式必传） - 2：月 - 3：年
	PeriodType *int32 `json:"period_type,omitempty"`

	// 订购周期数。（包周期模式必传） - period_type=2（周期类型为月），取值范围[1，9] - periodType=3（周期类型为年），取值范围[1，10]
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 是否自动续订。 - 1，自动续订 - 0，不自动续订  默认值为“0”
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	// 是否自动支付，下单订购后，是否自动从客户的华为云账户中支付，而不需要客户手动去进行支付。 - 1：是（会自动选择折扣和优惠券进行优惠，然后自动从客户华为云账户中支付），自动支付失败后会生成订单成功(该订单应付金额是优惠后金额)、但订单状态为“待支付”，等待客户手动支付(手动支付时，客户还可以修改系统自动选择的折扣和优惠券) - 0：否（需要客户手动去支付，客户可以选择折扣和优惠券。）  默认值为“0”
	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`

	Network *NetworkInfoCreate `json:"network"`

	// 云堡垒机实例是否支持IPV6。  默认值为“false”。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 企业项目ID。  默认值为“0”。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 附加磁盘大小。单位TB  > 说明： 附加磁盘和规格自带磁盘大小合起来不能超过300TB。
	AttachDiskSize *int32 `json:"attach_disk_size,omitempty"`

	// 标签信息。
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o CreateInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceBody", string(data)}, " ")
}
