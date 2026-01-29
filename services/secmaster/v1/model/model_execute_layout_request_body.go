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

type ExecuteLayoutRequestBody struct {

	// 对布局的操作 VERIFY 校验导入布局压缩包文件 IMPORT  导入布局压缩包文件 EXPORT 导出布局文件
	Action *def.MultiPart `json:"action"`

	// 操作布局的id列表
	Ids *def.MultiPart `json:"ids,omitempty"`

	// 上传的文件
	UploadFile *def.FilePart `json:"uploadFile,omitempty"`
}

func (o ExecuteLayoutRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteLayoutRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteLayoutRequestBody", string(data)}, " ")
}

func (o *ExecuteLayoutRequestBody) UnmarshalJSON(b []byte) error {
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

type ExecuteLayoutRequestBodyAction struct {
	value string
}

type ExecuteLayoutRequestBodyActionEnum struct {
	VERIFY ExecuteLayoutRequestBodyAction
	IMPORT ExecuteLayoutRequestBodyAction
	EXPORT ExecuteLayoutRequestBodyAction
}

func GetExecuteLayoutRequestBodyActionEnum() ExecuteLayoutRequestBodyActionEnum {
	return ExecuteLayoutRequestBodyActionEnum{
		VERIFY: ExecuteLayoutRequestBodyAction{
			value: "VERIFY",
		},
		IMPORT: ExecuteLayoutRequestBodyAction{
			value: "IMPORT",
		},
		EXPORT: ExecuteLayoutRequestBodyAction{
			value: "EXPORT",
		},
	}
}

func (c ExecuteLayoutRequestBodyAction) Value() string {
	return c.value
}

func (c ExecuteLayoutRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteLayoutRequestBodyAction) UnmarshalJSON(b []byte) error {
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
