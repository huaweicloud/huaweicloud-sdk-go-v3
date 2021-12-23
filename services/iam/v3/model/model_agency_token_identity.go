package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
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

func (c AgencyTokenIdentityMethods) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgencyTokenIdentityMethods) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
