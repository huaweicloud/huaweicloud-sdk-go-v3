package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriptionProductRequest Request Object
type ListSubscriptionProductRequest struct {

	// 用户当前语言环境， zh-cn简体中文，en-us英文环境
	XLanguage string `json:"X-Language"`
}

func (o ListSubscriptionProductRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionProductRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriptionProductRequest", string(data)}, " ")
}
