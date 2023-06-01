package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type QuotaRecord struct {

	// 记录ID。
	Id *string `json:"id,omitempty"`

	// 操作员的账号名称。
	Operator *string `json:"operator,omitempty"`

	// 操作类型。 10：发放额度11：回收额度
	OperationType *string `json:"operation_type,omitempty"`

	// 云经销商的代金券额度ID。 即华为云总经销商给云经销商发放代金券额度时，产生的云经销商的代金券额度ID，或者从云经销商回收代金券额度时，云经销商的代金券额度ID。
	QuotaId *string `json:"quota_id,omitempty"`

	// 父额度ID。 这即华为云总经销商给云经销商发放代金券额度时，华为云总经销商的额度ID，或者从云经销商回收代金券额度时，回收的华为云总经销商的额度ID。
	ParentQuotaId *string `json:"parent_quota_id,omitempty"`

	// 发放回收的金额。 取值大于0且精确到小数点后2位，单位：元。
	Amount *decimal.Decimal `json:"amount,omitempty"`

	// 操作时间，UTC时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。
	OperationTime *string `json:"operation_time,omitempty"`

	// 操作结果。 0：成功-1：失败
	Result *string `json:"result,omitempty"`

	// 云经销商的账号名。
	IndirectPartnerAccountName *string `json:"indirect_partner_account_name,omitempty"`

	// 云经销商ID。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`

	// 云经销商的公司名称。
	IndirectPartnerName *string `json:"indirect_partner_name,omitempty"`

	// 备注。
	Remark *string `json:"remark,omitempty"`
}

func (o QuotaRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaRecord struct{}"
	}

	return strings.Join([]string{"QuotaRecord", string(data)}, " ")
}
