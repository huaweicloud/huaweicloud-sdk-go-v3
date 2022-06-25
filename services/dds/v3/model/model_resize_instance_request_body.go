package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeInstanceRequestBody struct {
	Resize *ResizeInstanceOption `json:"resize"`

	// 变更包年包月实例规格时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。 - 对于降低规格场景，该字段无效。 - 对于扩大规格场景：   - true，表示自动从账户中支付。   - false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o ResizeInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeInstanceRequestBody", string(data)}, " ")
}
