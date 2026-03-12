package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDryRunConfigRequest Request Object
type ShowDryRunConfigRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 根的唯一标识符（ID）。
	RootId string `json:"root_id"`

	// 试运行策略的类型名称，service_control_policy：服务控制策略。
	PolicyType ShowDryRunConfigRequestPolicyType `json:"policy_type"`
}

func (o ShowDryRunConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDryRunConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowDryRunConfigRequest", string(data)}, " ")
}

type ShowDryRunConfigRequestPolicyType struct {
	value string
}

type ShowDryRunConfigRequestPolicyTypeEnum struct {
	SERVICE_CONTROL_POLICY ShowDryRunConfigRequestPolicyType
}

func GetShowDryRunConfigRequestPolicyTypeEnum() ShowDryRunConfigRequestPolicyTypeEnum {
	return ShowDryRunConfigRequestPolicyTypeEnum{
		SERVICE_CONTROL_POLICY: ShowDryRunConfigRequestPolicyType{
			value: "service_control_policy",
		},
	}
}

func (c ShowDryRunConfigRequestPolicyType) Value() string {
	return c.value
}

func (c ShowDryRunConfigRequestPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDryRunConfigRequestPolicyType) UnmarshalJSON(b []byte) error {
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
