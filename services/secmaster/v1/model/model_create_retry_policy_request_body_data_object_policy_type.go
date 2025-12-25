package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRetryPolicyRequestBodyDataObjectPolicyType 阻断类型
type CreateRetryPolicyRequestBodyDataObjectPolicyType struct {

	// 阻断类型：User Name/Source Ip/Domain Name
	PolicyType CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyType `json:"policy_type"`
}

func (o CreateRetryPolicyRequestBodyDataObjectPolicyType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetryPolicyRequestBodyDataObjectPolicyType struct{}"
	}

	return strings.Join([]string{"CreateRetryPolicyRequestBodyDataObjectPolicyType", string(data)}, " ")
}

type CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyType struct {
	value string
}

type CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyTypeEnum struct {
	USER_NAME   CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyType
	SOURCE_IP   CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyType
	DOMAIN_NAME CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyType
}

func GetCreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyTypeEnum() CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyTypeEnum {
	return CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyTypeEnum{
		USER_NAME: CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyType{
			value: "User Name",
		},
		SOURCE_IP: CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyType{
			value: "Source Ip",
		},
		DOMAIN_NAME: CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyType{
			value: "Domain Name",
		},
	}
}

func (c CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyType) Value() string {
	return c.value
}

func (c CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRetryPolicyRequestBodyDataObjectPolicyTypePolicyType) UnmarshalJSON(b []byte) error {
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
