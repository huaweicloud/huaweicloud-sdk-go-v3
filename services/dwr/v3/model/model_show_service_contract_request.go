package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowServiceContractRequest Request Object
type ShowServiceContractRequest struct {

	// 服务协议名称 使用公共Action的免责声明协议: use_public_action_privacy_statement 注册公共Action的免责声明协议: register_public_action_privacy_statement
	Type ShowServiceContractRequestType `json:"type"`
}

func (o ShowServiceContractRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServiceContractRequest struct{}"
	}

	return strings.Join([]string{"ShowServiceContractRequest", string(data)}, " ")
}

type ShowServiceContractRequestType struct {
	value string
}

type ShowServiceContractRequestTypeEnum struct {
	USE_PUBLIC_ACTION_PRIVACY_STATEMENT      ShowServiceContractRequestType
	REGISTER_PUBLIC_ACTION_PRIVACY_STATEMENT ShowServiceContractRequestType
}

func GetShowServiceContractRequestTypeEnum() ShowServiceContractRequestTypeEnum {
	return ShowServiceContractRequestTypeEnum{
		USE_PUBLIC_ACTION_PRIVACY_STATEMENT: ShowServiceContractRequestType{
			value: "use_public_action_privacy_statement",
		},
		REGISTER_PUBLIC_ACTION_PRIVACY_STATEMENT: ShowServiceContractRequestType{
			value: "register_public_action_privacy_statement",
		},
	}
}

func (c ShowServiceContractRequestType) Value() string {
	return c.value
}

func (c ShowServiceContractRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowServiceContractRequestType) UnmarshalJSON(b []byte) error {
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
