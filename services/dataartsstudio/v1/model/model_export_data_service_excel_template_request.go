package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportDataServiceExcelTemplateRequest Request Object
type ExportDataServiceExcelTemplateRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ExportDataServiceExcelTemplateRequestDlmType `json:"Dlm-Type,omitempty"`
}

func (o ExportDataServiceExcelTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDataServiceExcelTemplateRequest struct{}"
	}

	return strings.Join([]string{"ExportDataServiceExcelTemplateRequest", string(data)}, " ")
}

type ExportDataServiceExcelTemplateRequestDlmType struct {
	value string
}

type ExportDataServiceExcelTemplateRequestDlmTypeEnum struct {
	SHARED    ExportDataServiceExcelTemplateRequestDlmType
	EXCLUSIVE ExportDataServiceExcelTemplateRequestDlmType
}

func GetExportDataServiceExcelTemplateRequestDlmTypeEnum() ExportDataServiceExcelTemplateRequestDlmTypeEnum {
	return ExportDataServiceExcelTemplateRequestDlmTypeEnum{
		SHARED: ExportDataServiceExcelTemplateRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ExportDataServiceExcelTemplateRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ExportDataServiceExcelTemplateRequestDlmType) Value() string {
	return c.value
}

func (c ExportDataServiceExcelTemplateRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportDataServiceExcelTemplateRequestDlmType) UnmarshalJSON(b []byte) error {
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
