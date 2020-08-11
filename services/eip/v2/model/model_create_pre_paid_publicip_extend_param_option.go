/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 创建包周期弹性公网IP的订单信息
type CreatePrePaidPublicipExtendParamOption struct {
	// 功能说明：付费方式（预付费、按需付费；预付费，即包周期付费）  取值范围：  prePaid -预付费，即包年包月；  postPaid-后付费，即按需付费；  后付费的场景下，extendParam的其他字段都会被忽略。
	ChargeMode string `json:"charge_mode,omitempty"`
	// 功能说明：订购资源的周期类型（包年、包月等）  取值范围：  month-月  year-年  约束：如果用包周期共享带宽创建时（即携带共享带宽id创建弹性公网IP）此字段可不填。付费方式是预付费且不是使用共享带宽创建IP时，该字段必选；  使用共享带宽创建IP时，带宽资源到期时间与IP的到期时间相同。
	PeriodType string `json:"period_type,omitempty"`
	// 功能说明：订购周期数  取值范围：(后续会随运营策略变化)  period_type为month时，为[1,9]  period_type为year时，为[1,1]  约束：同period_type约束。
	PeriodNum int32 `json:"period_num,omitempty"`
	// 功能说明：是否自动续订  取值范围：  false：不自动续订  true：自动续订  约束：到期后，默认自动续订1个月（自动续订时间后续可能会变化），详情可联系客服咨询。
	IsAutoRenew bool `json:"is_auto_renew,omitempty"`
	// 功能说明：下单订购后，是否自动从客户的账户中支付  取值范围：  true：自动支付，从账户余额自动扣费  false：只提交订单不支付，需要客户手动去支付  约束：自动支付时，只能使用账户的现金支付；如果要使用代金券，请选择不自动支付，然后在用户费用中心，选择代金券支付。
	IsAutoPay bool `json:"is_auto_pay,omitempty"`
}
