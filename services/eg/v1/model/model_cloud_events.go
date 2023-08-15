package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloudEvents CloudEvents事件格式定义，参考https://github.com/cloudevents/spec/blob/v1.0.1/spec.md
type CloudEvents struct {

	// 事件唯一标识串，同一个事件来源里必须唯一
	Id string `json:"id"`

	// 事件来源上下文标识串，source+id可以唯一确定一个事件。采用URI-Reference格式，参考https://tools.ietf.org/html/rfc3986#section-4.1
	Source string `json:"source"`

	// CloudEvents协议版本，格式为major.minor
	Specversion string `json:"specversion"`

	// 事件类型
	Type string `json:"type"`

	// 事件内容格式，采用MIME格式，遵循RFC2046，参考https://tools.ietf.org/html/rfc2046
	Datacontenttype *string `json:"datacontenttype,omitempty"`

	// 事件内容模型定义的URI，遵循RFC3986，参考https://tools.ietf.org/html/rfc3986#section-4.3
	Dataschema *string `json:"dataschema,omitempty"`

	// 事件的负载内容，采用datacontenttype字段指定的格式，内容字段遵循dataschema字段的描述
	Data *interface{} `json:"data,omitempty"`

	// 事件发生UTC日期时间，相同来源的事件格式相同，遵循RFC3339，格式需满足2018-04-05T17:31:00Z，参考https://tools.ietf.org/html/rfc3339
	Time *string `json:"time,omitempty"`

	// 事件发生的主题或对象，用以标识哪个具体对象发生了当前事件
	Subject *string `json:"subject,omitempty"`
}

func (o CloudEvents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudEvents struct{}"
	}

	return strings.Join([]string{"CloudEvents", string(data)}, " ")
}
