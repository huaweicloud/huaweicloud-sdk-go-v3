package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SubTemplate 子模版数据结构
type SubTemplate struct {

	// 模板子类型，只支持以下5种类型：sms,dingding,wechat,webhook,email
	SubType SubTemplateSubType `json:"sub_type"`

	// 子模版正文，$符号后所跟变量仅支持以下变量，根据不同告警类型（关键词告警和sql告警），所支持的变量亦不相同。 目前两种告警类型有共同变量如下：告警级别：${event_severity};发生时间：${starts_at};告警源：$event.metadata.resource_provider;资源类型：$event.metadata.resource_type;资源标识：${resources};统计类型：关键词统计;表达式：$event.annotations.condition_expression;当前值: $event.annotations.current_value;统计周期：$event.annotations.frequency; 关键词告警特有变量：查询时间：$event.annotations.results[0].time;查询日志：$event.annotations.results[0].raw_results; sql告警特有变量：日志组/流名称：$event.annotations.results[0].resource_id;查询语句：$event.annotations.results[0].sql;查询时间：$event.annotations.results[0].time;查询URL：$event.annotations.results[0].url;查询日志：$event.annotations.results[0].raw_results; 变量后面的分号\";\"为英文符号，必须添加，否则模板会出现替换失败的情况
	Content string `json:"content"`

	// 邮件主题,只有sub_type=email时生效
	Topic *string `json:"topic,omitempty"`

	// **参数解释：**  当消息模板类型为webhook时生效，决定该消息的渲染方式。 **取值范围：**  - HTML - JSON
	SendType *string `json:"sendType,omitempty"`

	// **参数解释：**  消息模板的适用版本。 **取值范围：**   v1：标识为LTS的消息模板。
	Version *string `json:"version,omitempty"`
}

func (o SubTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTemplate struct{}"
	}

	return strings.Join([]string{"SubTemplate", string(data)}, " ")
}

type SubTemplateSubType struct {
	value string
}

type SubTemplateSubTypeEnum struct {
	SMS      SubTemplateSubType
	DINGDING SubTemplateSubType
	WECHAT   SubTemplateSubType
	WEBHOOK  SubTemplateSubType
	EMAIL    SubTemplateSubType
}

func GetSubTemplateSubTypeEnum() SubTemplateSubTypeEnum {
	return SubTemplateSubTypeEnum{
		SMS: SubTemplateSubType{
			value: "sms",
		},
		DINGDING: SubTemplateSubType{
			value: "dingding",
		},
		WECHAT: SubTemplateSubType{
			value: "wechat",
		},
		WEBHOOK: SubTemplateSubType{
			value: "webhook",
		},
		EMAIL: SubTemplateSubType{
			value: "email",
		},
	}
}

func (c SubTemplateSubType) Value() string {
	return c.value
}

func (c SubTemplateSubType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubTemplateSubType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
