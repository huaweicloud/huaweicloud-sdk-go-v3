package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAlertRuleSimulationRequestBody struct {

	// 数据管道 ID。Pipe ID.
	PipeId string `json:"pipe_id"`

	// 查询语句。Query.
	Query string `json:"query"`

	// 查询语法，SQL。Query type. SQL.
	QueryType *CreateAlertRuleSimulationRequestBodyQueryType `json:"query_type,omitempty"`

	// 开始时间。Start time.
	From int64 `json:"from"`

	// 结束时间。End time.
	To int64 `json:"to"`

	// 告警分组。Event grouping.
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// 告警触发规则。Alert triggers.
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
