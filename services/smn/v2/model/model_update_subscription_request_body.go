package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSubscriptionRequestBody struct {

	// 订阅者备注。订阅者备注的最大长度为128byte。
	Remark *string `json:"remark,omitempty"`

	// 订阅终端收到的验证码。
	VerificationCode *string `json:"verification_code,omitempty"`
}

func (o UpdateSubscriptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionRequestBody", string(data)}, " ")
}
