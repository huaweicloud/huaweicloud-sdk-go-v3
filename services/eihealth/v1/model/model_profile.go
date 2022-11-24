package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

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
