package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateSubTemplate struct {

	// 模板子类型，只支持以下类型：sms，dingding（钉钉），wechat（微信），webhook，email
	SubType UpdateSubTemplateSubType `json:"sub_type"`

	// 子模板正文，$符号后所跟变量仅支持以下变量，根据不同告警类型（关键词告警和sql告警），所支持的变量亦不相同。目前两种告警类型有共同变量如下：  告警级别：${event_severity}; 发生时间：${starts_at}; 告警源：$event.metadata.resource_provider; 资源类型：$event.metadata.resource_type; 资源标识：${resources}; 统计类型：关键词统计; 表达式：$event.annotations.condition_expression; 当前值: $event.annotations.current_value; 统计周期：$event.annotations.frequency; 关键词告警特有变量： 查询时间：$event.annotations.results[0].time; 查询日志：$event.annotations.results[0].raw_results; sql告警特有变量： 日志组/流名称：$event.annotations.results[0].resource_id; 查询语句：$event.annotations.results[0].sql; 查询时间：$event.annotations.results[0].time; 查询URL：$event.annotations.results[0].url; 查询日志：$event.annotations.results[0].raw_results;   >变量后面的分号\";\"为英文符号，必须添加，否则模板会出现替换失败的情况。
	Content string `json:"content"`

	// 邮件主题,只有sub_type=email时生效
	Topic *string `json:"topic,omitempty"`
}

func (o UpdateSubTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubTemplate struct{}"
	}

	return strings.Join([]string{"UpdateSubTemplate", string(data)}, " ")
}

type UpdateSubTemplateSubType struct {
	value string
}

type UpdateSubTemplateSubTypeEnum struct {
	SMS      UpdateSubTemplateSubType
	DINGDING UpdateSubTemplateSubType
	WECHAT   UpdateSubTemplateSubType
	WEBHOOK  UpdateSubTemplateSubType
	EMAIL    UpdateSubTemplateSubType
	VOICE    UpdateSubTemplateSubType
}

func GetUpdateSubTemplateSubTypeEnum() UpdateSubTemplateSubTypeEnum {
	return UpdateSubTemplateSubTypeEnum{
		SMS: UpdateSubTemplateSubType{
			value: "sms",
		},
		DINGDING: UpdateSubTemplateSubType{
			value: "dingding",
		},
		WECHAT: UpdateSubTemplateSubType{
			value: "wechat",
		},
		WEBHOOK: UpdateSubTemplateSubType{
			value: "webhook",
		},
		EMAIL: UpdateSubTemplateSubType{
			value: "email",
		},
		VOICE: UpdateSubTemplateSubType{
			value: "voice",
		},
	}
}

func (c UpdateSubTemplateSubType) Value() string {
	return c.value
}

func (c UpdateSubTemplateSubType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSubTemplateSubType) UnmarshalJSON(b []byte) error {
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
