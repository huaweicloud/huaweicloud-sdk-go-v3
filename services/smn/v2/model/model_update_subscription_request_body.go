package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSubscriptionRequestBody struct {

	// 订阅者备注。订阅者备注的最大长度为128byte。
	Remark string `json:"remark"`
}

func (o UpdateSubscriptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionRequestBody", string(data)}, " ")
}
