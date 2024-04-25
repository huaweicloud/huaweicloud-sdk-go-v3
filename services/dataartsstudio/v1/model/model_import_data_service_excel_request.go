package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImportDataServiceExcelRequest Request Object
type ImportDataServiceExcelRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ImportDataServiceExcelRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），此处为导入文件，取值为multipart/form-data。
	ContentType string `json:"Content-Type"`

	Body *ImportDataServiceExcelRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportDataServiceExcelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataServiceExcelRequest struct{}"
	}

	return strings.Join([]string{"ImportDataServiceExcelRequest", string(data)}, " ")
}

type ImportDataServiceExcelRequestDlmType struct {
	value string
}

type ImportDataServiceExcelRequestDlmTypeEnum struct {
	SHARED    ImportDataServiceExcelRequestDlmType
	EXCLUSIVE ImportDataServiceExcelRequestDlmType
}

func GetImportDataServiceExcelRequestDlmTypeEnum() ImportDataServiceExcelRequestDlmTypeEnum {
	return ImportDataServiceExcelRequestDlmTypeEnum{
		SHARED: ImportDataServiceExcelRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ImportDataServiceExcelRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ImportDataServiceExcelRequestDlmType) Value() string {
	return c.value
}

func (c ImportDataServiceExcelRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportDataServiceExcelRequestDlmType) UnmarshalJSON(b []byte) error {
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
