package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPartnerAccountChangeRecordsRequest Request Object
type ListPartnerAccountChangeRecordsRequest struct {

	// 账户类型。BALANCE_TYPE_DEBIT：现金账户,BALANCE_TYPE_CREDIT：信用账户
	BalanceType string `json:"balance_type"`

	// 交易类型。RECHARGE：充值,DEDEUCT：消费,REFUND：退款,RFROZEN：冻结,TRANS：转账，余额和保证金互换（老商务模式，当前已无保证金账户）,ADJUST：调账（华为核销等）,BEADJUST：经销商拨款,BERETRIEVE：经销商回收,BEUNBIND：解绑/关联模式切换导致的回收,BONUSCONVERT：奖励金转换（老商务模式，当前已无奖励金账户）,TRADE_MODE_TRANSFER：交易模式变更, 此参数不携带或携带值为空时，不作为筛选条件。
	TradeType *string `json:"trade_type,omitempty"`

	// 查询收支明细的开始日期。 说明： 东八区时间，格式为YYYY-MM-DD，如“2017-10-21”。默认值为一年前的当天日期。
	TradeTimeBegin *string `json:"trade_time_begin,omitempty"`

	// 查询收支明细的结束日期。 说明： 东八区时间，格式为YYYY-MM-DD，如“2017-12-21”。默认值为当前日期。
	TradeTimeEnd *string `json:"trade_time_end,omitempty"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的数量，默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。  说明： 华为云总经销商（一级经销商）查询云经销商（二级经销商）的收支明细时，需携带此参数；除此之外，此参数不作处理。否则只能查询自身的收支明细。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListPartnerAccountChangeRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartnerAccountChangeRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListPartnerAccountChangeRecordsRequest", string(data)}, " ")
}
