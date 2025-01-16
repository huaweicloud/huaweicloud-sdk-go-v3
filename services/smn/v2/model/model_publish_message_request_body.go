package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublishMessageRequestBody struct {

	// 消息标题，给邮箱订阅者发送邮件时作为邮件主题，长度不能超过512个字节。
	Subject *string `json:"subject,omitempty"`

	// 发送的消息。消息体必须是UTF-8编码的字符串，大小至多256KB。如果订阅者是手机号码，短信长度限制为490字，超出则可能被运营商拦截。短信内容不能包含“[]”或者“【】”符号。 说明： 三种消息发送方式：message、message_structure、message_template_name  至少设置其中一个，如果同时设置，生效的优先级为 message_structure > message_template_name > message。
	Message *string `json:"message,omitempty"`

	// Json格式的字符串。支持“email”、“sms”、“http”、“https”、“functiongraph”、“functionstage”、“dingding”、“wechat”、“feishu”、“welink”。必须设置默认的消息“default”，当匹配不到消息协议时，按“default”中的内容发送。其中，钉钉、微信、飞书、welink协议类型的消息需指定msgType字段；钉钉，微信和飞书机器人协议支持msgType为text（纯文本）和markdown（MD）格式消息，welink和红版welink机器人类型暂仅支持msgType为text的纯文本消息。   钉钉机器人协议支持通过at字段实现@群组成员。当您需要@群成员时，可以在isAtAll字段中传入布尔值，表示是否需要@群组内所有人。您可以在atMobiles字段中传入需要@的人的手机号列表，或在atUserIds字段中传入需要@的钉钉userID列表。当您使用atMobiles字段或atUserIds字段时，需要在消息内容中同时传入@对应手机号或userID的信息。展示效果请参考钉钉官方文档。  说明： 三种消息发送方式：message、message_structure、message_template_name  至少设置其中一个，如果同时设置，生效的优先级为 message_structure > message_template_name > message。
	MessageStructure *string `json:"message_structure,omitempty"`

	// 消息模板名称，可通过[查询消息模板列表](ListMessageTemplates.xml)获取名称。  说明： 三种消息发送方式：message、message_structure、message_template_name  至少设置其中一个，如果同时设置，生效的优先级为 message_structure > message_template_name > message。
	MessageTemplateName *string `json:"message_template_name,omitempty"`

	// tag以及替换tag的参数组成的字典。消息模板中的标签对应的值。使用消息模板方式的消息发布必须携带该参数。字典中的key为消息模板中的参数名称，不超过21个字符。字典中的value为消息模板中的参数被替换后的值，不超过1KB。
	Tags map[string]string `json:"tags,omitempty"`

	// 指消息在SMN系统内部的最长存留时间。超过该存留时间，系统将不再发送该消息。单位是s，变量默认值是3600s，即一小时。值为正整数且小于等于3600*24。
	TimeToLive *string `json:"time_to_live,omitempty"`

	// 消息属性列表
	MessageAttributes *[]MessageAttribute `json:"message_attributes,omitempty"`

	// 语言，发送出去的消息中SMN附加系统内容的语言，若没传入，使用账号的语言。取值范围是该局点支持的语言，比如：zh-cn,en-us等
	Locale *string `json:"locale,omitempty"`
}

func (o PublishMessageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishMessageRequestBody struct{}"
	}

	return strings.Join([]string{"PublishMessageRequestBody", string(data)}, " ")
}
