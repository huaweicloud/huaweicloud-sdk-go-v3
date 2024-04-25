package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportDataServiceExcelRequest Request Object
type ExportDataServiceExcelRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ExportDataServiceExcelRequestDlmType `json:"Dlm-Type,omitempty"`

	// API导出ID列表。
	Body *[]string `json:"body,omitempty"`
}

func (o ExportDataServiceExcelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDataServiceExcelRequest struct{}"
	}

	return strings.Join([]string{"ExportDataServiceExcelRequest", string(data)}, " ")
}

type ExportDataServiceExcelRequestDlmType struct {
	value string
}

type ExportDataServiceExcelRequestDlmTypeEnum struct {
	SHARED    ExportDataServiceExcelRequestDlmType
	EXCLUSIVE ExportDataServiceExcelRequestDlmType
}

func GetExportDataServiceExcelRequestDlmTypeEnum() ExportDataServiceExcelRequestDlmTypeEnum {
	return ExportDataServiceExcelRequestDlmTypeEnum{
		SHARED: ExportDataServiceExcelRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ExportDataServiceExcelRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ExportDataServiceExcelRequestDlmType) Value() string {
	return c.value
}

func (c ExportDataServiceExcelRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportDataServiceExcelRequestDlmType) UnmarshalJSON(b []byte) error {
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
