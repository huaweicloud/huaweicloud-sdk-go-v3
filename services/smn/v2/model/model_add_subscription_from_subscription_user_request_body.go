package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddSubscriptionFromSubscriptionUserRequestBody struct {

	// 订阅用户ID列表。
	Ids []string `json:"ids"`
}

func (o AddSubscriptionFromSubscriptionUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubscriptionFromSubscriptionUserRequestBody struct{}"
	}

	return strings.Join([]string{"AddSubscriptionFromSubscriptionUserRequestBody", string(data)}, " ")
}
