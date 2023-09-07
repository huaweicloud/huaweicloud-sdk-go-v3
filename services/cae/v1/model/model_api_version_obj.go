package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApiVersionObj API版本，固定值“v1”，该值不可修改。
type ApiVersionObj struct {
	value string
}

type ApiVersionObjEnum struct {
	V1 ApiVersionObj
}

func GetApiVersionObjEnum() ApiVersionObjEnum {
	return ApiVersionObjEnum{
		V1: ApiVersionObj{
			value: "v1",
		},
	}
}

func (c ApiVersionObj) Value() string {
	return c.value
}

func (c ApiVersionObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiVersionObj) UnmarshalJSON(b []byte) error {
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
