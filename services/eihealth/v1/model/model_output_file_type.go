package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OutputFileType struct {
	value string
}

type OutputFileTypeEnum struct {
	TXT OutputFileType
	VCF OutputFileType
	CSV OutputFileType
}

func GetOutputFileTypeEnum() OutputFileTypeEnum {
	return OutputFileTypeEnum{
		TXT: OutputFileType{
			value: "txt",
		},
		VCF: OutputFileType{
			value: "vcf",
		},
		CSV: OutputFileType{
			value: "csv",
		},
	}
}

func (c OutputFileType) Value() string {
	return c.value
}

func (c OutputFileType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OutputFileType) UnmarshalJSON(b []byte) error {
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
