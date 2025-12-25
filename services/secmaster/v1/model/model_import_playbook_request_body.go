package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

	"strings"
)

type ImportPlaybookRequestBody struct {

	// 导入文件
	UploadFile *def.FilePart `json:"uploadFile"`

	// 剧本文件上传模式 import 剧本上传导入 verify 剧本上传校验
	UploadModel *def.MultiPart `json:"upload_model,omitempty"`
}

func (o ImportPlaybookRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportPlaybookRequestBody struct{}"
	}

	return strings.Join([]string{"ImportPlaybookRequestBody", string(data)}, " ")
}

func (o *ImportPlaybookRequestBody) UnmarshalJSON(b []byte) error {
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

type ImportPlaybookRequestBodyUploadModel struct {
	value string
}

type ImportPlaybookRequestBodyUploadModelEnum struct {
	VERIFY ImportPlaybookRequestBodyUploadModel
	IMPORT ImportPlaybookRequestBodyUploadModel
}

func GetImportPlaybookRequestBodyUploadModelEnum() ImportPlaybookRequestBodyUploadModelEnum {
	return ImportPlaybookRequestBodyUploadModelEnum{
		VERIFY: ImportPlaybookRequestBodyUploadModel{
			value: "verify",
		},
		IMPORT: ImportPlaybookRequestBodyUploadModel{
			value: "import",
		},
	}
}

func (c ImportPlaybookRequestBodyUploadModel) Value() string {
	return c.value
}

func (c ImportPlaybookRequestBodyUploadModel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportPlaybookRequestBodyUploadModel) UnmarshalJSON(b []byte) error {
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
