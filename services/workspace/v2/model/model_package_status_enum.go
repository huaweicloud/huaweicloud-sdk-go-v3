package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PackageStatusEnum 技能包状态枚举。
type PackageStatusEnum struct {
	value string
}

type PackageStatusEnumEnum struct {
	UPLOAD    PackageStatusEnum
	PUBLISHED PackageStatusEnum
	DISABLED  PackageStatusEnum
}

func GetPackageStatusEnumEnum() PackageStatusEnumEnum {
	return PackageStatusEnumEnum{
		UPLOAD: PackageStatusEnum{
			value: "UPLOAD",
		},
		PUBLISHED: PackageStatusEnum{
			value: "PUBLISHED",
		},
		DISABLED: PackageStatusEnum{
			value: "DISABLED",
		},
	}
}

func (c PackageStatusEnum) Value() string {
	return c.value
}

func (c PackageStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PackageStatusEnum) UnmarshalJSON(b []byte) error {
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
