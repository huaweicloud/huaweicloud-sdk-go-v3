package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UploadDbObjectTemplateRequest Request Object
type UploadDbObjectTemplateRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *UploadDbObjectTemplateRequestXLanguage `json:"X-Language,omitempty"`

	// 文件模板支持数据同步级别，不填默认为table表级。 - database：库级 - table：表级 - column：列级
	FileImportDbLevel *UploadDbObjectTemplateRequestFileImportDbLevel `json:"file_import_db_level,omitempty"`

	Body *UploadDbObjectTemplateRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadDbObjectTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadDbObjectTemplateRequest struct{}"
	}

	return strings.Join([]string{"UploadDbObjectTemplateRequest", string(data)}, " ")
}

type UploadDbObjectTemplateRequestXLanguage struct {
	value string
}

type UploadDbObjectTemplateRequestXLanguageEnum struct {
	EN_US UploadDbObjectTemplateRequestXLanguage
	ZH_CN UploadDbObjectTemplateRequestXLanguage
}

func GetUploadDbObjectTemplateRequestXLanguageEnum() UploadDbObjectTemplateRequestXLanguageEnum {
	return UploadDbObjectTemplateRequestXLanguageEnum{
		EN_US: UploadDbObjectTemplateRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UploadDbObjectTemplateRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UploadDbObjectTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c UploadDbObjectTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UploadDbObjectTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type UploadDbObjectTemplateRequestFileImportDbLevel struct {
	value string
}

type UploadDbObjectTemplateRequestFileImportDbLevelEnum struct {
	DATABASE UploadDbObjectTemplateRequestFileImportDbLevel
	TABLE    UploadDbObjectTemplateRequestFileImportDbLevel
	COLUMN   UploadDbObjectTemplateRequestFileImportDbLevel
}

func GetUploadDbObjectTemplateRequestFileImportDbLevelEnum() UploadDbObjectTemplateRequestFileImportDbLevelEnum {
	return UploadDbObjectTemplateRequestFileImportDbLevelEnum{
		DATABASE: UploadDbObjectTemplateRequestFileImportDbLevel{
			value: "database",
		},
		TABLE: UploadDbObjectTemplateRequestFileImportDbLevel{
			value: "table",
		},
		COLUMN: UploadDbObjectTemplateRequestFileImportDbLevel{
			value: "column",
		},
	}
}

func (c UploadDbObjectTemplateRequestFileImportDbLevel) Value() string {
	return c.value
}

func (c UploadDbObjectTemplateRequestFileImportDbLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UploadDbObjectTemplateRequestFileImportDbLevel) UnmarshalJSON(b []byte) error {
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
