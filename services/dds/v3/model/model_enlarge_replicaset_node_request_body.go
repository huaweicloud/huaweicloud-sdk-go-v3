package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnlargeReplicasetNodeRequestBody struct {

	// 副本集节点扩容个数，副本集有3个节点时，可以扩容2/4个节点，副本集有5个节点时，只能扩容2个
	Num int32 `json:"num"`

	// 变更包年包月实例规格时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。 - 对于降低规格场景，该字段无效。 - 对于扩大规格场景：   - true，表示自动从账户中支付。   - false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o EnlargeReplicasetNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeReplicasetNodeRequestBody struct{}"
	}

	return strings.Join([]string{"EnlargeReplicasetNodeRequestBody", string(data)}, " ")
}
