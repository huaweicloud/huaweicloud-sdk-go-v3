package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAgencyCredentialRequest Request Object
type ShowAgencyCredentialRequest struct {

	// 委托类型：TABLE_SERVICE_TRUST-表服务委托。
	AgencyType *ShowAgencyCredentialRequestAgencyType `json:"agency_type,omitempty"`
}

func (o ShowAgencyCredentialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyCredentialRequest struct{}"
	}

	return strings.Join([]string{"ShowAgencyCredentialRequest", string(data)}, " ")
}

type ShowAgencyCredentialRequestAgencyType struct {
	value string
}

type ShowAgencyCredentialRequestAgencyTypeEnum struct {
	TABLE_SERVICE_TRUST ShowAgencyCredentialRequestAgencyType
}

func GetShowAgencyCredentialRequestAgencyTypeEnum() ShowAgencyCredentialRequestAgencyTypeEnum {
	return ShowAgencyCredentialRequestAgencyTypeEnum{
		TABLE_SERVICE_TRUST: ShowAgencyCredentialRequestAgencyType{
			value: "TABLE_SERVICE_TRUST",
		},
	}
}

func (c ShowAgencyCredentialRequestAgencyType) Value() string {
	return c.value
}

func (c ShowAgencyCredentialRequestAgencyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAgencyCredentialRequestAgencyType) UnmarshalJSON(b []byte) error {
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
