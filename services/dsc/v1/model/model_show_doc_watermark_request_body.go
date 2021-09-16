package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"strings"
)

type ShowDocWatermarkRequestBody struct {
	// 待提取水印的文档类型
	DocType *def.MultiPart `json:"doc_type"`

	// 解密文件的密码， 最大支持长度256。如果Office文档有读密码或域控的权限密码，请输入读密码，或者有读权限的域控密码。
	FilePassword *def.MultiPart `json:"file_password,omitempty"`

	// 上传要提取水印的文档
	File *def.FilePart `json:"file"`
}

func (o ShowDocWatermarkRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDocWatermarkRequestBody struct{}"
	}

	return strings.Join([]string{"ShowDocWatermarkRequestBody", string(data)}, " ")
}

type ShowDocWatermarkRequestBodyDocType struct {
	value string
}

type ShowDocWatermarkRequestBodyDocTypeEnum struct {
	WORD  ShowDocWatermarkRequestBodyDocType
	EXCEL ShowDocWatermarkRequestBodyDocType
	PDF   ShowDocWatermarkRequestBodyDocType
	PPT   ShowDocWatermarkRequestBodyDocType
}

func GetShowDocWatermarkRequestBodyDocTypeEnum() ShowDocWatermarkRequestBodyDocTypeEnum {
	return ShowDocWatermarkRequestBodyDocTypeEnum{
		WORD: ShowDocWatermarkRequestBodyDocType{
			value: "WORD",
		},
		EXCEL: ShowDocWatermarkRequestBodyDocType{
			value: "EXCEL",
		},
		PDF: ShowDocWatermarkRequestBodyDocType{
			value: "PDF",
		},
		PPT: ShowDocWatermarkRequestBodyDocType{
			value: "PPT",
		},
	}
}

func (c ShowDocWatermarkRequestBodyDocType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowDocWatermarkRequestBodyDocType) UnmarshalJSON(b []byte) error {
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
