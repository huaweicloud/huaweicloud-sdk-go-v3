package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AgencyTokenIdentity
type AgencyTokenIdentity struct {

	// token的获取方式，该字段内容为[\"assume_role\"]。
	Methods []AgencyTokenIdentityMethods `json:"methods"`

	AssumeRole *AgencyTokenAssumerole `json:"assume_role"`
}

func (o AgencyTokenIdentity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyTokenIdentity struct{}"
	}

	return strings.Join([]string{"AgencyTokenIdentity", string(data)}, " ")
}

type AgencyTokenIdentityMethods struct {
	value string
}

type AgencyTokenIdentityMethodsEnum struct {
	ASSUME_ROLE AgencyTokenIdentityMethods
}

func GetAgencyTokenIdentityMethodsEnum() AgencyTokenIdentityMethodsEnum {
	return AgencyTokenIdentityMethodsEnum{
		ASSUME_ROLE: AgencyTokenIdentityMethods{
			value: "assume_role",
		},
	}
}

func (c AgencyTokenIdentityMethods) Value() string {
	return c.value
}

func (c AgencyTokenIdentityMethods) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgencyTokenIdentityMethods) UnmarshalJSON(b []byte) error {
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
