package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAiGuardRuleStatusRequestBody 更新规则状态请求
type UpdateAiGuardRuleStatusRequestBody struct {

	// 规则状态，可选：enable(启用), disable(禁用)
	Status UpdateAiGuardRuleStatusRequestBodyStatus `json:"status"`
}

func (o UpdateAiGuardRuleStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiGuardRuleStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAiGuardRuleStatusRequestBody", string(data)}, " ")
}

type UpdateAiGuardRuleStatusRequestBodyStatus struct {
	value string
}

type UpdateAiGuardRuleStatusRequestBodyStatusEnum struct {
	ENABLE  UpdateAiGuardRuleStatusRequestBodyStatus
	DISABLE UpdateAiGuardRuleStatusRequestBodyStatus
}

func GetUpdateAiGuardRuleStatusRequestBodyStatusEnum() UpdateAiGuardRuleStatusRequestBodyStatusEnum {
	return UpdateAiGuardRuleStatusRequestBodyStatusEnum{
		ENABLE: UpdateAiGuardRuleStatusRequestBodyStatus{
			value: "enable",
		},
		DISABLE: UpdateAiGuardRuleStatusRequestBodyStatus{
			value: "disable",
		},
	}
}

func (c UpdateAiGuardRuleStatusRequestBodyStatus) Value() string {
	return c.value
}

func (c UpdateAiGuardRuleStatusRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAiGuardRuleStatusRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
