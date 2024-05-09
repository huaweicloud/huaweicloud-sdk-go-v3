package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SearchType struct {
	value string
}

type SearchTypeEnum struct {
	ECFP_4   SearchType
	SCAFFOLD SearchType
}

func GetSearchTypeEnum() SearchTypeEnum {
	return SearchTypeEnum{
		ECFP_4: SearchType{
			value: "ECFP_4",
		},
		SCAFFOLD: SearchType{
			value: "SCAFFOLD",
		},
	}
}

func (c SearchType) Value() string {
	return c.value
}

func (c SearchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchType) UnmarshalJSON(b []byte) error {
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
