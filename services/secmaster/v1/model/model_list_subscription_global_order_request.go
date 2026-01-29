package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriptionGlobalOrderRequest Request Object
type ListSubscriptionGlobalOrderRequest struct {

	// 用户当前语言环境， zh-cn简体中文，en-us英文环境
	XLanguage string `json:"X-Language"`
}

func (o ListSubscriptionGlobalOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionGlobalOrderRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriptionGlobalOrderRequest", string(data)}, " ")
}
