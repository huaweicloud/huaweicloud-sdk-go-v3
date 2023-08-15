package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckLigandDifferenceMethod 差异计算方法：RMSD、EMD
type CheckLigandDifferenceMethod struct {
	value string
}

type CheckLigandDifferenceMethodEnum struct {
	RMSD CheckLigandDifferenceMethod
	EMD  CheckLigandDifferenceMethod
}

func GetCheckLigandDifferenceMethodEnum() CheckLigandDifferenceMethodEnum {
	return CheckLigandDifferenceMethodEnum{
		RMSD: CheckLigandDifferenceMethod{
			value: "RMSD",
		},
		EMD: CheckLigandDifferenceMethod{
			value: "EMD",
		},
	}
}

func (c CheckLigandDifferenceMethod) Value() string {
	return c.value
}

func (c CheckLigandDifferenceMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckLigandDifferenceMethod) UnmarshalJSON(b []byte) error {
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
