package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddSubscriptionRequestBody struct {

	// 不同协议对应不同的endpoint（接受消息的接入点）。 目前支持的协议包括：  “email”：邮件传输协议，endpoint为邮箱地址。  “sms”：短信传输协议，endpoint为手机号码。  “functionstage”：FunctionGraph（函数）传输协议，endpoint为一个函数。  “functiongraph”：FunctionGraph（工作流）传输协议，endpoint为由一组函数编排成的工作流。  “http”、“https”：HTTP/HTTPS传输协议，endpoint为URL。  “callnotify”：语音通知传输协议，endpoint为手机号码。  “wechat”：微信群机器人传输协议。  “dingding”：钉钉群机器人传输协议。  “feishu”：飞书群机器人传输协议。  “welink”：welink群机器人传输协议。  “dingTalkBot”：个人钉钉传输协议。
	Protocol string `json:"protocol"`

	// 说明：  http协议，接入点必须以“http://”开头，且必须为公网的IPV4地址、IPV6地址或域名地址。不支持使用内网地址，包括但不限于IPv4私有地址、IPv6本地地址、回环地址等。需确保该地址可通过公网路由访问，避免因防火墙、NAT或DNS配置导致不可达。  https协议，接入点必须以“https://”开头，且必须为公网的IPV4地址、IPV6地址或域名地址。不支持使用内网地址，包括但不限于IPv4私有地址、IPv6本地地址、回环地址等。需确保该地址可通过公网路由访问，避免因防火墙、NAT或DNS配置导致不可达。  email协议，接入点必须是邮件地址。  sms协议，接入点必须是一个电话号码。  functionstage协议，接入点必须是一个函数。  functiongraph协议，接入点必须是一个函数工作流。  callnotify协议，接入点必须是一个电话号码。  dingding协议，接入点必须是一个钉钉自定义群机器人的地址，或添加了钉钉企业内部机器人的群组openConversationId。添加钉钉企业内部机器人对应的群组openConversationId时，该字段不能以“http://”或“https://”为开头。  wechat协议，接入点必须是一个微信群机器人的地址。  feishu协议，接入点必须是一个飞书群机器人的地址。  welink协议，接入点必须是一个welink的群号。  dingTalkBot协议，接入点必须是一个钉钉企业用户的userId。
	Endpoint string `json:"endpoint"`

	// 备注。最大支持128字节，约42个中文，必须是UTF-8编码的字符串，否则无法正常显示中文。
	Remark *string `json:"remark,omitempty"`

	Extension *SubscriptionExtension `json:"extension,omitempty"`

	// 当需要批量创建订阅时，需要传入该字段。SMN支持一次最多创建50个订阅。
	Subscriptions *[]BatchAddSubscriptionsRequestBody `json:"subscriptions,omitempty"`
}

func (o AddSubscriptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubscriptionRequestBody struct{}"
	}

	return strings.Join([]string{"AddSubscriptionRequestBody", string(data)}, " ")
}
