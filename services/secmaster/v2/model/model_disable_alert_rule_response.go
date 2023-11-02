package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DisableAlertRuleResponse Response Object
type DisableAlertRuleResponse struct {

	// 告警规则 ID。Alert rule ID.
	RuleId *string `json:"rule_id,omitempty"`

	// 启用状态，启用、停用。Status, enabled, disabled.
	Status *DisableAlertRuleResponseStatus `json:"status,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableAlertRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableAlertRuleResponse struct{}"
	}

	return strings.Join([]string{"DisableAlertRuleResponse", string(data)}, " ")
}

type DisableAlertRuleResponseStatus struct {
	value string
}

type DisableAlertRuleResponseStatusEnum struct {
	ENABLED  DisableAlertRuleResponseStatus
	DISABLED DisableAlertRuleResponseStatus
}

func GetDisableAlertRuleResponseStatusEnum() DisableAlertRuleResponseStatusEnum {
	return DisableAlertRuleResponseStatusEnum{
		ENABLED: DisableAlertRuleResponseStatus{
			value: "ENABLED",
		},
		DISABLED: DisableAlertRuleResponseStatus{
			value: "DISABLED",
		},
	}
}

func (c DisableAlertRuleResponseStatus) Value() string {
	return c.value
}

func (c DisableAlertRuleResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisableAlertRuleResponseStatus) UnmarshalJSON(b []byte) error {
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
