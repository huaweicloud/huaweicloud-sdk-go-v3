package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdjustRecordV3 struct {

	// 客户账号ID。
	CustomerId *string `json:"customer_id,omitempty"`

	// 客户名称。
	CustomerName *string `json:"customer_name,omitempty"`

	// 调账类型。 SOURCE_OPERATION_BEADJUST：拨款SOURCE_OPERATION_BERETRIEVE：回收SOURCE_OPERATION_BEUNBIND：解绑回收
	OperationType *string `json:"operation_type,omitempty"`

	// 调账的总金额。
	Amount *string `json:"amount,omitempty"`

	// 币种。 CNY：人民币
	Currency *string `json:"currency,omitempty"`

	// 使用场景。
	ApplyScene *string `json:"apply_scene,omitempty"`

	// 调账操作的时间。 UTC时间，格式为：2016-03-28T14:45:38Z
	OperationTime *string `json:"operation_time,omitempty"`

	// 调账单位。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 事务ID。
	TransId *string `json:"trans_id,omitempty"`
}

func (o AdjustRecordV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdjustRecordV3 struct{}"
	}

	return strings.Join([]string{"AdjustRecordV3", string(data)}, " ")
}
