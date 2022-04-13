package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExpandInstanceNodeRequestBody struct {
	// 新增加的节点数量。

	Num int32 `json:"num"`
	// 创建包周期实例时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。 - true，表示自动从账户中支付。 - false，表示手动从账户中支付，默认为该方式。

	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o ExpandInstanceNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandInstanceNodeRequestBody struct{}"
	}

	return strings.Join([]string{"ExpandInstanceNodeRequestBody", string(data)}, " ")
}
