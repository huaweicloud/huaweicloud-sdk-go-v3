package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddSubscriptionRequestBody struct {

	// 不同协议对应不同的endpoint（接受消息的接入点）。 目前支持的协议包括：  “email”：邮件传输协议，endpoint为邮箱地址。  “sms”：短信传输协议，endpoint为手机号码。  “functionstage”：FunctionGraph（函数）传输协议，endpoint为一个函数。  “functiongraph”：FunctionGraph（工作流）传输协议，endpoint为由一组函数编排成的工作流。  “http”、“https”：HTTP/HTTPS传输协议，endpoint为URL。  “callnotify”：语音通知传输协议，endpoint为手机号码。  “wechat”：微信群机器人传输协议。  “dingding”：钉钉群机器人传输协议。  “feishu”：飞书群机器人传输协议。  “welink”：welink群机器人传输协议。
	Protocol string `json:"protocol"`

	// 说明：  http协议，接入点必须以“http://”开头。  https协议，接入点必须以“https://”开头。  email协议，接入点必须是邮件地址。  sms协议，接入点必须是一个电话号码。  functionstage协议，接入点必须是一个函数。  functiongraph协议，接入点必须是一个函数工作流。  dms协议，接入点必须是一个消息队列。  application协议，接入点必须是一个应用平台的设备终端。  callnotify协议，接入点必须是一个电话号码。  dingding协议，接入点必须是一个钉钉群机器人的地址。  wechat协议，接入点必须是一个微信群机器人的地址。  feishu协议，接入点必须是一个飞书群机器人的地址。  welink协议，接入点必须是一个welink的群号。
	Endpoint string `json:"endpoint"`

	// 备注。最大支持128字节，约42个中文，必须是UTF-8编码的字符串，否则无法正常显示中文。
	Remark *string `json:"remark,omitempty"`

	Extension *SubscriptionExtension `json:"extension,omitempty"`
}

func (o AddSubscriptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubscriptionRequestBody struct{}"
	}

	return strings.Join([]string{"AddSubscriptionRequestBody", string(data)}, " ")
}
