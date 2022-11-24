package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会议推送参数配置请求。
type WebHookConfigRequest struct {

	// 企业ID。按企业注册回调时需要填写。
	CorpId *string `json:"corpId,omitempty"`

	// SP ID。多租户场景下，按SP注册回调时需要填写。
	SpId *string `json:"spId,omitempty"`

	// 订阅者ID。
	SubscriberId string `json:"subscriberId"`

	// 订阅者秘钥。
	SubscriberKey string `json:"subscriberKey"`

	// 订阅url。 > 必须使用HTTPS。
	Url string `json:"url"`
}

func (o WebHookConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebHookConfigRequest struct{}"
	}

	return strings.Join([]string{"WebHookConfigRequest", string(data)}, " ")
}
