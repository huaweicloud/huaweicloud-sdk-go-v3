package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 变更至目标规格的请求信息。
type ResizeFlavorReq struct {

	// 变更至新规格的资源规格编码。
	SpecCode string `json:"spec_code"`

	// 实例默认一个组，此时不需要传入该参数。当使用组功能创建多个组时, 需要传入需要规格变更的对应组的group_id。
	GroupId *string `json:"group_id,omitempty"`

	// 变更包年包月实例规格时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。true，表示自动从账户中支付。false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o ResizeFlavorReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeFlavorReq struct{}"
	}

	return strings.Join([]string{"ResizeFlavorReq", string(data)}, " ")
}
