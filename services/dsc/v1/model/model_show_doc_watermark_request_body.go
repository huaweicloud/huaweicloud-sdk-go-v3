package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDocWatermarkRequestBody struct{}"
	}

	return strings.Join([]string{"ShowDocWatermarkRequestBody", string(data)}, " ")
}

func (o *ShowDocWatermarkRequestBody) UnmarshalJSON(b []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	t := reflect.TypeOf(o).Elem()
	v := reflect.ValueOf(o).Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		jsonName := strings.Split(jsonTag, ",")[0]
		if m[jsonName] == nil && strings.Contains(jsonTag, "omitempty") {
			continue
		}
		field := v.FieldByName(utils.UnderscoreToCamel(jsonName))
		switch v.Field(i).Interface().(type) {
		case *def.FilePart:
			filePath := m[jsonName].(string)
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(def.NewFilePart(file)))
		case *def.MultiPart:
			field.Set(reflect.ValueOf(def.NewMultiPart(m[jsonName])))
		default:
			return errors.New(fmt.Sprintf("unmarshal %s failed", m[jsonName]))
		}
	}
	return nil
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

func (c ShowDocWatermarkRequestBodyDocType) Value() string {
	return c.value
}

func (c ShowDocWatermarkRequestBodyDocType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDocWatermarkRequestBodyDocType) UnmarshalJSON(b []byte) error {
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
