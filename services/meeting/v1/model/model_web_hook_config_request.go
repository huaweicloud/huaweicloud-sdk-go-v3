package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会议推送参数配置请求体
type WebHookConfigRequest struct {

	// 企业ID，与spId二者必填一个
	CorpId *string `json:"corpId,omitempty" xml:"corpId"`

	// sp管理员ID，与corpId二者必填一个
	SpId *string `json:"spId,omitempty" xml:"spId"`

	// 订阅者ID
	SubscriberId string `json:"subscriberId" xml:"subscriberId"`

	// 订阅者秘钥
	SubscriberKey string `json:"subscriberKey" xml:"subscriberKey"`

	// 订阅url，建议使用HTTPS
	Url string `json:"url" xml:"url"`
}

func (o WebHookConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebHookConfigRequest struct{}"
	}

	return strings.Join([]string{"WebHookConfigRequest", string(data)}, " ")
}
