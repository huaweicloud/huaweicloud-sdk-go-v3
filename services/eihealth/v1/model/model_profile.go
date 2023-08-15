package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Profile struct {
	value string
}

type ProfileEnum struct {
	PY3 Profile
}

func GetProfileEnum() ProfileEnum {
	return ProfileEnum{
		PY3: Profile{
			value: "PY3",
		},
	}
}

func (c Profile) Value() string {
	return c.value
}

func (c Profile) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Profile) UnmarshalJSON(b []byte) error {
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
