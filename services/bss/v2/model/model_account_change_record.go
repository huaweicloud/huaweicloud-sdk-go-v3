package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccountChangeRecord struct {

	// 收支明细流水号。
	AccountChangeId *string `json:"account_change_id,omitempty"`

	// 交易详细类型。 SOURCE_OPERAIION_ADJUST_CREDIT：调账(信用额度调整)SOURCE_OPERAIION_RECHARGE：充值SOURCE_OPERAIION_DEDEUCT：消费(包年/包月)SOURCE_OPERAIION_MANUALFROZE：冻结(人工冻结)SOURCE_OPERAIION_MANUALUNFROZE：冻结(人工解冻)SOURCE_OPERAIION_MANUALCLEARFROZEN：冻结(人工清零)SOURCE_OPERAIION_TRANS_TO_BALANCE：转账(保证金转余额)SOURCE_OPERATION_BEADJUST：伙伴拨款SOURCE_OPERATION_BEUNBIND：交易模式变更(切换/解除关联回收)SOURCE_OPERAIION_EXPIRECLEAR：过期清零SOURCE_OPERAIION_ONETIME：消费(一次性扣费)SOURCE_OPERAIION_REFUND：退款SOURCE_OPERAIION_UNFROZEN：退款(退款解冻)SOURCE_OPERAIION_CLEARFROZEN：退款(退款清零)SOURCE_OPERAIION_ADJUST：调账(余额调整)SOURCE_OPERAIION_USAGE：消费(按需)SOURCE_OPERAIION_WRITEOFF：消费(欠费还款)SOURCE_OPERAIION_UNSUBSCRIBE：退款SOURCE_OPERAIION_RFROZEN：退款(退款冻结)SOURCE_OPERAIION_TRANS_TO_FOREGIFT：转账(余额转保证金)SOURCE_OPERAIION_PRIZE：调账(赠送)SOURCE_OPERATION_BERETRIEVE：伙伴回收SOURCE_OPERAIION_PRECISIONCOMP：消费(精度补扣)SOURCE_OPERAIION_FREERESDEDUCT：消费(免费资源扣减)SOURCE_OPERAIION_MERGE：奖励金转换(合并)SOURCE_OPERAIION_CONVERT_BONUS：奖励金转换SOURCE_OPERAIION_RECHARGE_REBATE：充值(激励返点)SOURCE_OPERATION_COUPONCANCEL：优惠券回收SOURCE_OPERAIION_BILLREFUND：调账(华为核销)SOURCE_OPERATION_TRADEMODE_TRANSFER：交易模式变更(和伙伴关联)SOURCE_OPERATION_SYSTEM_FROZEN：系统冻结（购买标销合同的伙伴涉及该模式）SOURCE_OPERATION_SYSTEM_UNFROZEN：系统解冻（购买标销合同的伙伴涉及该模式）SOURCE_OPERATION_COUPON_QUOTA_TRANSFER：调账(兑换现金券额度)SOURCE_OPERATIION_RIDEDUCT：消费(预留实例)SOURCE_OPERATION_COUPON_QUOTA_RECLAIM：代金券回收
	TradeDetailType *string `json:"trade_detail_type,omitempty"`

	// 交易时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2016-03-28T14:45:38Z”。
	TradeTime *string `json:"trade_time,omitempty"`

	// 交易ID/订单ID。
	TradeId *string `json:"trade_id,omitempty"`

	// 变更金额，单位为元。
	ChangeAmount *string `json:"change_amount,omitempty"`

	// 变更后余额，单位为元。
	BalanceAfterChange *string `json:"balance_after_change,omitempty"`

	// 收支类型。 1：收入2：支出
	Type *string `json:"type,omitempty"`
}

func (o AccountChangeRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountChangeRecord struct{}"
	}

	return strings.Join([]string{"AccountChangeRecord", string(data)}, " ")
}
