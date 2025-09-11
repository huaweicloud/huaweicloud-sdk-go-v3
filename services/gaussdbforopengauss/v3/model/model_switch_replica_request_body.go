package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchReplicaRequestBody struct {

	// **参数解释**: 创建包周期实例时可指定，表示是否自动从账户中支付，该字段不影响自动续订的支付方式。 **约束限制**: 不涉及。 **取值范围**: - true：表示自动从账户中支付。 - false：表示手动从账户中支付，默认为该支付方式。  **默认取值**: 不涉及。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o SwitchReplicaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchReplicaRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchReplicaRequestBody", string(data)}, " ")
}
