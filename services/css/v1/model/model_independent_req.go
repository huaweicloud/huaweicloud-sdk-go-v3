package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IndependentReq struct {
	Type *IndependentBodyReq `json:"type"`

	// 是否自动支付。下单订购后，是否自动从客户的华为云账户中支付，而不需要客户手动去进行支付。
	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`
}

func (o IndependentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndependentReq struct{}"
	}

	return strings.Join([]string{"IndependentReq", string(data)}, " ")
}
