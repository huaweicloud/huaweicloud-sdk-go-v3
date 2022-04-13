package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会议推送参数配置请求体
type WebHookConfigRequest struct {
	// 企业ID，与spId二者必填一个

	CorpId *string `json:"corpId,omitempty"`
	// sp管理员ID，与corpId二者必填一个

	SpId *string `json:"spId,omitempty"`
	// 订阅者ID

	SubscriberId string `json:"subscriberId"`
	// 订阅者秘钥

	SubscriberKey string `json:"subscriberKey"`
	// 订阅url，建议使用HTTPS

	Url string `json:"url"`
}

func (o WebHookConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebHookConfigRequest struct{}"
	}

	return strings.Join([]string{"WebHookConfigRequest", string(data)}, " ")
}
