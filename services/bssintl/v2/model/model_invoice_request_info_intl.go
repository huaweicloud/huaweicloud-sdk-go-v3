package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type InvoiceRequestInfoIntl struct {

	// 请求ID。
	RequestId *string `json:"requestId,omitempty"`

	// 驳回原因。
	CancelReason *string `json:"cancelReason,omitempty"`

	// 开票类型。 0：个人1：企业
	TitleType *int32 `json:"titleType,omitempty"`

	// 渠道类型。 0：华为云
	ChannelType *int32 `json:"channelType,omitempty"`

	// 发票种类。 0：增值税专用发票1：增值税普通发票
	InvoiceType *int32 `json:"invoiceType,omitempty"`

	// 发票抬头。
	InvoiceTitle *string `json:"invoiceTitle,omitempty"`

	// 已开票金额（美元） 。
	InvoiceAmount *decimal.Decimal `json:"invoiceAmount,omitempty"`

	// 开票方式。 0：账期1：到账2：订单
	InvoiceMethod *int32 `json:"invoiceMethod,omitempty"`

	// 发票类别。 0：税票1：商票
	InvoiceClass *int32 `json:"invoiceClass,omitempty"`

	// 开票状态。 0：草稿1：待审核4：等待导出发票文件5：等待发票文件回填6：等待邮寄确认7：等待回执确认8：完成9：已退票11：等待驳回审核13：退票待审核14：待退票状态回填15：退票失败
	InvoiceState *int32 `json:"invoiceState,omitempty"`

	// 发票申请人员。
	ApplyOpera *string `json:"applyOpera,omitempty"`

	AddressInfo *PostAddressInfoIntl `json:"addressInfo,omitempty"`

	// 申请时间（UTC时间）。
	ApplyTime *string `json:"applyTime,omitempty"`

	// 发票类型。 0：纸质票
	InvoiceMode *string `json:"invoiceMode,omitempty"`

	// 电子发票寄送地。
	Email *string `json:"email,omitempty"`

	// 申请类型。 0：开票申请1：退票申请2：正向开票已退票
	RequestMode *string `json:"requestMode,omitempty"`

	// 退票时的原申请ID。
	SrcRequestId *string `json:"srcRequestId,omitempty"`

	// 签约主体ID。
	SalesId *string `json:"salesId,omitempty"`

	// 发票号码。
	InvoiceNo *string `json:"invoiceNo,omitempty"`

	// 交易类型。 3：结算信用卡扣减4：结算未结清开票5：先开票后到款6：BP月结开票7：充值开票8：包年/包月在线支付开票10：普通提现开票
	TradeType *int32 `json:"tradeType,omitempty"`

	// 发票账期。
	BillCycle *string `json:"billCycle,omitempty"`

	// 税务信息列表，参见表4。
	TaxList *[]TaxInfo `json:"taxList,omitempty"`
}

func (o InvoiceRequestInfoIntl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvoiceRequestInfoIntl struct{}"
	}

	return strings.Join([]string{"InvoiceRequestInfoIntl", string(data)}, " ")
}
