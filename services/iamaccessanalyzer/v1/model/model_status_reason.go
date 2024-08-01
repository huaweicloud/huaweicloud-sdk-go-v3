package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StatusReason 提供有关分析器当前状态的更多详细信息。
type StatusReason struct {

	// 分析器当前状态的原因。
	Code StatusReasonCode `json:"code"`

	// 分析器当前状态的详细原因。
	Details *string `json:"details,omitempty"`
}

func (o StatusReason) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusReason struct{}"
	}

	return strings.Join([]string{"StatusReason", string(data)}, " ")
}

type StatusReasonCode struct {
	value string
}

type StatusReasonCodeEnum struct {
	DELEGATED_ADMINISTRATOR_DEREGISTERED  StatusReasonCode
	TRUSTED_SERVICE_DISABLED              StatusReasonCode
	INTERNAL_ERROR                        StatusReasonCode
	ORGANIZATION_DELETED                  StatusReasonCode
	SERVICE_LINKED_AGENCY_CREATION_FAILED StatusReasonCode
}

func GetStatusReasonCodeEnum() StatusReasonCodeEnum {
	return StatusReasonCodeEnum{
		DELEGATED_ADMINISTRATOR_DEREGISTERED: StatusReasonCode{
			value: "delegated_administrator_deregistered",
		},
		TRUSTED_SERVICE_DISABLED: StatusReasonCode{
			value: "trusted_service_disabled",
		},
		INTERNAL_ERROR: StatusReasonCode{
			value: "internal_error",
		},
		ORGANIZATION_DELETED: StatusReasonCode{
			value: "organization_deleted",
		},
		SERVICE_LINKED_AGENCY_CREATION_FAILED: StatusReasonCode{
			value: "service_linked_agency_creation_failed",
		},
	}
}

func (c StatusReasonCode) Value() string {
	return c.value
}

func (c StatusReasonCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StatusReasonCode) UnmarshalJSON(b []byte) error {
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
