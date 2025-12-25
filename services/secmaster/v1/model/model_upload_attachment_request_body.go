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

type UploadAttachmentRequestBody struct {

	// 固定枚举，当前仅支持user_task/用户任务
	FileType *def.MultiPart `json:"fileType,omitempty"`

	// 导入的流程文件
	UploadFile *def.FilePart `json:"uploadFile"`
}

func (o UploadAttachmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAttachmentRequestBody struct{}"
	}

	return strings.Join([]string{"UploadAttachmentRequestBody", string(data)}, " ")
}

func (o *UploadAttachmentRequestBody) UnmarshalJSON(b []byte) error {
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

type UploadAttachmentRequestBodyFileType struct {
	value string
}

type UploadAttachmentRequestBodyFileTypeEnum struct {
	USER_TASK UploadAttachmentRequestBodyFileType
}

func GetUploadAttachmentRequestBodyFileTypeEnum() UploadAttachmentRequestBodyFileTypeEnum {
	return UploadAttachmentRequestBodyFileTypeEnum{
		USER_TASK: UploadAttachmentRequestBodyFileType{
			value: "user_task",
		},
	}
}

func (c UploadAttachmentRequestBodyFileType) Value() string {
	return c.value
}

func (c UploadAttachmentRequestBodyFileType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UploadAttachmentRequestBodyFileType) UnmarshalJSON(b []byte) error {
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
