package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CouponRecordV2 struct {

	// 该记录的ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 操作类型。 1：发放2：手动回收3：解绑自动回收
	OperationType *string `json:"operation_type,omitempty" xml:"operation_type"`

	// 额度ID。
	QuotaId *string `json:"quota_id,omitempty" xml:"quota_id"`

	// 额度类型。 0：代金券额度1：现金券额度
	QuotaType *int32 `json:"quota_type,omitempty" xml:"quota_type"`

	// 代金券ID。
	CouponId *string `json:"coupon_id,omitempty" xml:"coupon_id"`

	// 客户账号ID。
	CustomerId *string `json:"customer_id,omitempty" xml:"customer_id"`

	// 操作的面额值。单位：元。 发放时，等于面额值；回收时，指每次回收的具体值。
	OperationAmount *float64 `json:"operation_amount,omitempty" xml:"operation_amount"`

	// 操作时间。
	OperationTime *string `json:"operation_time,omitempty" xml:"operation_time"`

	// 操作结果。 0：成功-1：失败
	Result *string `json:"result,omitempty" xml:"result"`

	// 操作记录中的备注。
	Remark *string `json:"remark,omitempty" xml:"remark"`
}

func (o CouponRecordV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CouponRecordV2 struct{}"
	}

	return strings.Join([]string{"CouponRecordV2", string(data)}, " ")
}
