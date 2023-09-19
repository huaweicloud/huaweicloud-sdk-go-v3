package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandClusterResponse Response Object
type ExpandClusterResponse struct {

	// 请求操作结果，succeeded为操作成功，failed为操作失败。非包周期节点组扩容请求下发成功时，会包含该字段且内容为success。
	Result *string `json:"result,omitempty"`

	// 订单ID。对包周期节点组进行扩容时，会返回本次扩容产生的订单ID，需要客户到订单支付页面进行自主支付才能真正触发扩容。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandClusterResponse struct{}"
	}

	return strings.Join([]string{"ExpandClusterResponse", string(data)}, " ")
}
