package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAlertRuleSimulationRequestBody struct {

	// 告警分组
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// 开始时间
	From int64 `json:"from"`

	// 数据管道 ID
	PipeId string `json:"pipe_id"`

	// 查询语句
	Query string `json:"query"`

	// **参数解释**: 查询语法类型 - SQL查询 **约束限制**: 不涉及  **取值范围**: - SQL **默认值**:  SQL
	QueryType *CreateAlertRuleSimulationRequestBodyQueryType `json:"query_type,omitempty"`

	// 结束时间
	To int64 `json:"to"`

	// 告警触发规则
	Triggers []AlertRuleTrigger `json:"triggers"`
}

func (o CreateAlertRuleSimulationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertRuleSimulationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAlertRuleSimulationRequestBody", string(data)}, " ")
}

type CreateAlertRuleSimulationRequestBodyQueryType struct {
	value string
}

type CreateAlertRuleSimulationRequestBodyQueryTypeEnum struct {
	SQL CreateAlertRuleSimulationRequestBodyQueryType
}

func GetCreateAlertRuleSimulationRequestBodyQueryTypeEnum() CreateAlertRuleSimulationRequestBodyQueryTypeEnum {
	return CreateAlertRuleSimulationRequestBodyQueryTypeEnum{
		SQL: CreateAlertRuleSimulationRequestBodyQueryType{
			value: "SQL",
		},
	}
}

func (c CreateAlertRuleSimulationRequestBodyQueryType) Value() string {
	return c.value
}

func (c CreateAlertRuleSimulationRequestBodyQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAlertRuleSimulationRequestBodyQueryType) UnmarshalJSON(b []byte) error {
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
