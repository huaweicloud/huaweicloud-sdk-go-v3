package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OrderV5 struct {

	// 订单id
	OrderId *string `json:"order_id,omitempty"`

	// 订单状态:0:初始化; 1:待审核; 2:待退款; 3:处理中; 4:已取消; 5:已完成; 6:待支付; 7:补偿中; 8:待审批; 9:待确认; 10:待发货; 11:待收货; 12:待上门取货; 13:换新中; 14:待商家收货
	OrderStatus *int32 `json:"order_status,omitempty"`

	// 结果，SUCCESS:成功； FAIL：失败
	Result *string `json:"result,omitempty"`

	// result=FAIL时，必填，标识该订单失败原因
	ResultCode *string `json:"result_code,omitempty"`

	// 失败信息，和result_code结对出现
	ResultMsg *string `json:"result_msg,omitempty"`
}

func (o OrderV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderV5 struct{}"
	}

	return strings.Join([]string{"OrderV5", string(data)}, " ")
}
