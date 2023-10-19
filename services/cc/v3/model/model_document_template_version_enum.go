package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DocumentTemplateVersionEnum 文档模板版本。 - 2022.08.30 (2022.08.30)
type DocumentTemplateVersionEnum struct {
	value string
}

type DocumentTemplateVersionEnumEnum struct {
	E_2022_08_30 DocumentTemplateVersionEnum
}

func GetDocumentTemplateVersionEnumEnum() DocumentTemplateVersionEnumEnum {
	return DocumentTemplateVersionEnumEnum{
		E_2022_08_30: DocumentTemplateVersionEnum{
			value: "2022.08.30",
		},
	}
}

func (c DocumentTemplateVersionEnum) Value() string {
	return c.value
}

func (c DocumentTemplateVersionEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DocumentTemplateVersionEnum) UnmarshalJSON(b []byte) error {
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
