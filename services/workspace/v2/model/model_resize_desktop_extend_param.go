package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopExtendParam 桌面变更规格的扩展参数。
type ResizeDesktopExtendParam struct {

	// 是否自动付款，下单订购后，是否自动从客户的账户中支付，而不需要客户手动去进行支付。默认为false。
	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o ResizeDesktopExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopExtendParam struct{}"
	}

	return strings.Join([]string{"ResizeDesktopExtendParam", string(data)}, " ")
}
