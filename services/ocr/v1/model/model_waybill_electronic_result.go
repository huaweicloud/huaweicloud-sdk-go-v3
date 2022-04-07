package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WaybillElectronicResult struct {
	// 三段码。

	Code *string `json:"code,omitempty"`
	// 收件人姓名。

	ReceiverName *string `json:"receiver_name,omitempty"`
	// 收件人电话。

	ReceiverPhone *string `json:"receiver_phone,omitempty"`
	// 收件人地址。

	ReceiverAddress *string `json:"receiver_address,omitempty"`
	// 寄件人姓名。

	SenderName *string `json:"sender_name,omitempty"`
	// 寄件人电话。

	SenderPhone *string `json:"sender_phone,omitempty"`
	// 寄件人地址。

	SenderAddress *string `json:"sender_address,omitempty"`
	// 条形码运单号。

	WaybillNumber *string `json:"waybill_number,omitempty"`
	// 相关字段的置信度信息，取值范围0~1。 置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。

	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o WaybillElectronicResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WaybillElectronicResult struct{}"
	}

	return strings.Join([]string{"WaybillElectronicResult", string(data)}, " ")
}
