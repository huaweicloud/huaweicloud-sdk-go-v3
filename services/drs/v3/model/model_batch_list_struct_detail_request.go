package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchListStructDetailRequest struct {
	// 数据库支持迁移对象类型

	Type BatchListStructDetailRequestType `json:"type"`
	// 请求语言类型

	XLanguage BatchListStructDetailRequestXLanguage `json:"X-Language"`

	Body *BatchQueryJobReqPage `json:"body,omitempty"`
}

func (o BatchListStructDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListStructDetailRequest struct{}"
	}

	return strings.Join([]string{"BatchListStructDetailRequest", string(data)}, " ")
}

type BatchListStructDetailRequestType struct {
	value string
}

type BatchListStructDetailRequestTypeEnum struct {
	DATABASE        BatchListStructDetailRequestType
	SCHEMA          BatchListStructDetailRequestType
	TABLE           BatchListStructDetailRequestType
	VIEW            BatchListStructDetailRequestType
	PROCEDURE       BatchListStructDetailRequestType
	TRIGGER         BatchListStructDetailRequestType
	INDEX           BatchListStructDetailRequestType
	TABLE_INDEXS    BatchListStructDetailRequestType
	TABLE_STRUCTURE BatchListStructDetailRequestType
}

func GetBatchListStructDetailRequestTypeEnum() BatchListStructDetailRequestTypeEnum {
	return BatchListStructDetailRequestTypeEnum{
		DATABASE: BatchListStructDetailRequestType{
			value: "database",
		},
		SCHEMA: BatchListStructDetailRequestType{
			value: "schema",
		},
		TABLE: BatchListStructDetailRequestType{
			value: "table",
		},
		VIEW: BatchListStructDetailRequestType{
			value: "view",
		},
		PROCEDURE: BatchListStructDetailRequestType{
			value: "procedure",
		},
		TRIGGER: BatchListStructDetailRequestType{
			value: "trigger",
		},
		INDEX: BatchListStructDetailRequestType{
			value: "index",
		},
		TABLE_INDEXS: BatchListStructDetailRequestType{
			value: "table_indexs",
		},
		TABLE_STRUCTURE: BatchListStructDetailRequestType{
			value: "table_structure",
		},
	}
}

func (c BatchListStructDetailRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListStructDetailRequestType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type BatchListStructDetailRequestXLanguage struct {
	value string
}

type BatchListStructDetailRequestXLanguageEnum struct {
	EN_US BatchListStructDetailRequestXLanguage
	ZH_CN BatchListStructDetailRequestXLanguage
}

func GetBatchListStructDetailRequestXLanguageEnum() BatchListStructDetailRequestXLanguageEnum {
	return BatchListStructDetailRequestXLanguageEnum{
		EN_US: BatchListStructDetailRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchListStructDetailRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchListStructDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListStructDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
