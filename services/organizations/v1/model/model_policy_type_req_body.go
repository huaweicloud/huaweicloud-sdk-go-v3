package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyTypeReqBody 策略类型相关操作的请求体。
type PolicyTypeReqBody struct {

	// 策略类型的名称，service_control_policy服务控制策略；tag_policy：标签策略。
	PolicyType PolicyTypeReqBodyPolicyType `json:"policy_type"`

	// 根的唯一标识符（ID）。
	RootId string `json:"root_id"`
}

func (o PolicyTypeReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTypeReqBody struct{}"
	}

	return strings.Join([]string{"PolicyTypeReqBody", string(data)}, " ")
}

type PolicyTypeReqBodyPolicyType struct {
	value string
}

type PolicyTypeReqBodyPolicyTypeEnum struct {
	SERVICE_CONTROL_POLICY PolicyTypeReqBodyPolicyType
	TAG_POLICY             PolicyTypeReqBodyPolicyType
}

func GetPolicyTypeReqBodyPolicyTypeEnum() PolicyTypeReqBodyPolicyTypeEnum {
	return PolicyTypeReqBodyPolicyTypeEnum{
		SERVICE_CONTROL_POLICY: PolicyTypeReqBodyPolicyType{
			value: "service_control_policy",
		},
		TAG_POLICY: PolicyTypeReqBodyPolicyType{
			value: "tag_policy",
		},
	}
}

func (c PolicyTypeReqBodyPolicyType) Value() string {
	return c.value
}

func (c PolicyTypeReqBodyPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyTypeReqBodyPolicyType) UnmarshalJSON(b []byte) error {
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
