package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAlertRuleSimulationRequestBody struct {

	// pipe_id
	PipeId string `json:"pipe_id"`

	// query
	Query string `json:"query"`

	// query_type. SQL, CBSL.
	QueryType *CreateAlertRuleSimulationRequestBodyQueryType `json:"query_type,omitempty"`

	// from
	From int64 `json:"from"`

	// from
	To int64 `json:"to"`

	// event_grouping
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// triggers
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
	SQL  CreateAlertRuleSimulationRequestBodyQueryType
	CBSL CreateAlertRuleSimulationRequestBodyQueryType
}

func GetCreateAlertRuleSimulationRequestBodyQueryTypeEnum() CreateAlertRuleSimulationRequestBodyQueryTypeEnum {
	return CreateAlertRuleSimulationRequestBodyQueryTypeEnum{
		SQL: CreateAlertRuleSimulationRequestBodyQueryType{
			value: "SQL",
		},
		CBSL: CreateAlertRuleSimulationRequestBodyQueryType{
			value: "CBSL",
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
