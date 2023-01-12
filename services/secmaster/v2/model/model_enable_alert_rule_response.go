package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type EnableAlertRuleResponse struct {

	// rule_id
	RuleId *string `json:"rule_id,omitempty"`

	// status. ENABLED, DISABLED
	Status *EnableAlertRuleResponseStatus `json:"status,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableAlertRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAlertRuleResponse struct{}"
	}

	return strings.Join([]string{"EnableAlertRuleResponse", string(data)}, " ")
}

type EnableAlertRuleResponseStatus struct {
	value string
}

type EnableAlertRuleResponseStatusEnum struct {
	ENABLED  EnableAlertRuleResponseStatus
	DISABLED EnableAlertRuleResponseStatus
}

func GetEnableAlertRuleResponseStatusEnum() EnableAlertRuleResponseStatusEnum {
	return EnableAlertRuleResponseStatusEnum{
		ENABLED: EnableAlertRuleResponseStatus{
			value: "ENABLED",
		},
		DISABLED: EnableAlertRuleResponseStatus{
			value: "DISABLED",
		},
	}
}

func (c EnableAlertRuleResponseStatus) Value() string {
	return c.value
}

func (c EnableAlertRuleResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnableAlertRuleResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
