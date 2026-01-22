package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CheckDomainVerificationInfo struct {

	// 直播域名
	Domain string `json:"domain"`

	// 验证方式，DNS：DNS解析验证；FILE：文件验证
	VerifyType CheckDomainVerificationInfoVerifyType `json:"verify_type"`
}

func (o CheckDomainVerificationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDomainVerificationInfo struct{}"
	}

	return strings.Join([]string{"CheckDomainVerificationInfo", string(data)}, " ")
}

type CheckDomainVerificationInfoVerifyType struct {
	value string
}

type CheckDomainVerificationInfoVerifyTypeEnum struct {
	DNS  CheckDomainVerificationInfoVerifyType
	FILE CheckDomainVerificationInfoVerifyType
}

func GetCheckDomainVerificationInfoVerifyTypeEnum() CheckDomainVerificationInfoVerifyTypeEnum {
	return CheckDomainVerificationInfoVerifyTypeEnum{
		DNS: CheckDomainVerificationInfoVerifyType{
			value: "DNS",
		},
		FILE: CheckDomainVerificationInfoVerifyType{
			value: "FILE",
		},
	}
}

func (c CheckDomainVerificationInfoVerifyType) Value() string {
	return c.value
}

func (c CheckDomainVerificationInfoVerifyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckDomainVerificationInfoVerifyType) UnmarshalJSON(b []byte) error {
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
