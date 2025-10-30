package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubTemplateResBody struct {

	// **参数解释：**  消息模板的通知渠道。 **取值范围：**  - sms - dingding - wechat - webhook - email - voice - feishu - welink
	SubType *string `json:"sub_type,omitempty"`

	// **参数解释：**  对应通知渠道的消息模板正文，$符号后所跟变量仅支持以下变量，根据不同告警类型（关键词告警和sql告警），所支持的变量亦不相同。目前两种告警类型有共同变量如下： 告警级别：${event_severity}; 发生时间：${starts_at}; 告警源：$event.metadata.resource_provider; 资源类型：$event.metadata.resource_type; 资源标识：${resources}; 统计类型：关键词统计; 表达式：$event.annotations.condition_expression; 当前值: $event.annotations.current_value; 统计周期：$event.annotations.frequency; 关键词告警特有变量： 查询时间：$event.annotations.results[0].time; 查询日志：$event.annotations.results[0].raw_results; sql告警特有变量： 日志组/流名称：$event.annotations.results[0].resource_id; 查询语句：$event.annotations.results[0].sql; 查询时间：$event.annotations.results[0].time; 查询URL：$event.annotations.results[0].url; 查询日志：$event.annotations.results[0].raw_results;  告警级别：${event_severity}; 发生时间：${starts_at}; 告警源：$event.metadata.resource_provider; 资源类型：$event.metadata.resource_type; 资源标识：${resources}; 统计类型：关键词统计; 表达式：$event.annotations.condition_expression; 当前值: $event.annotations.current_value; 统计周期：$event.annotations.frequency; 关键词告警特有变量： 查询时间：$event.annotations.results[0].time; 查询日志：$event.annotations.results[0].raw_results; sql告警特有变量： 日志组/流名称：$event.annotations.results[0].resource_id; 查询语句：$event.annotations.results[0].sql; 查询时间：$event.annotations.results[0].time; 查询URL：$event.annotations.results[0].url; 查询日志：$event.annotations.results[0].raw_results;   >变量后面的分号\";\"为英文符号，必须添加，否则模板会出现替换失败的情况。  **取值范围：**  不涉及。
	Content *string `json:"content,omitempty"`

	// **参数解释：**  邮件主题。通知渠道为邮件时生效，即sub_type=email。 **取值范围：**  不涉及。
	Topic *string `json:"topic,omitempty"`

	// **参数解释：**  当消息模板类型为webhook时生效，决定该消息的渲染方式。 **取值范围：**  - HTML - JSON
	SendType *string `json:"sendType,omitempty"`

	// **参数解释：**  消息模板的适用版本。 **取值范围：**   v1：标识为LTS的消息模板。
	Version *string `json:"version,omitempty"`
}

func (o SubTemplateResBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTemplateResBody struct{}"
	}

	return strings.Join([]string{"SubTemplateResBody", string(data)}, " ")
}
