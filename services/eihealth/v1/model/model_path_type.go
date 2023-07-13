package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PathType struct {
	value string
}

type PathTypeEnum struct {
	FILE   PathType
	FOLDER PathType
}

func GetPathTypeEnum() PathTypeEnum {
	return PathTypeEnum{
		FILE: PathType{
			value: "FILE",
		},
		FOLDER: PathType{
			value: "FOLDER",
		},
	}
}

func (c PathType) Value() string {
	return c.value
}

func (c PathType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PathType) UnmarshalJSON(b []byte) error {
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
