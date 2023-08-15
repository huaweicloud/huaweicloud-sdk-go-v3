package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AcceptServiceContractRequest Request Object
type AcceptServiceContractRequest struct {

	// 服务协议名称 使用公共Action的免责声明协议: use_public_action_privacy_statement 注册公共Action的免责声明协议: register_public_action_privacy_statement
	Type AcceptServiceContractRequestType `json:"type"`
}

func (o AcceptServiceContractRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptServiceContractRequest struct{}"
	}

	return strings.Join([]string{"AcceptServiceContractRequest", string(data)}, " ")
}

type AcceptServiceContractRequestType struct {
	value string
}

type AcceptServiceContractRequestTypeEnum struct {
	USE_PUBLIC_ACTION_PRIVACY_STATEMENT      AcceptServiceContractRequestType
	REGISTER_PUBLIC_ACTION_PRIVACY_STATEMENT AcceptServiceContractRequestType
}

func GetAcceptServiceContractRequestTypeEnum() AcceptServiceContractRequestTypeEnum {
	return AcceptServiceContractRequestTypeEnum{
		USE_PUBLIC_ACTION_PRIVACY_STATEMENT: AcceptServiceContractRequestType{
			value: "use_public_action_privacy_statement",
		},
		REGISTER_PUBLIC_ACTION_PRIVACY_STATEMENT: AcceptServiceContractRequestType{
			value: "register_public_action_privacy_statement",
		},
	}
}

func (c AcceptServiceContractRequestType) Value() string {
	return c.value
}

func (c AcceptServiceContractRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AcceptServiceContractRequestType) UnmarshalJSON(b []byte) error {
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
