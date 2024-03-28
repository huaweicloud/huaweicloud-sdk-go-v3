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

type ImportApiDefinitionsAsyncRequestBody struct {

	// 是否创建新分组
	IsCreateGroup *def.MultiPart `json:"is_create_group,omitempty"`

	// API分组编号，当is_create_group=false时为必填
	GroupId *def.MultiPart `json:"group_id,omitempty"`

	// 扩展信息导入模式 - merge：当扩展信息定义冲突时，merge保留原有扩展信息 - override：当扩展信息定义冲突时，override会覆盖原有扩展信息
	ExtendMode *def.MultiPart `json:"extend_mode,omitempty"`

	// 是否开启简易导入模式
	SimpleMode *def.MultiPart `json:"simple_mode,omitempty"`

	// 是否开启Mock后端
	MockMode *def.MultiPart `json:"mock_mode,omitempty"`

	// 导入模式 - merge：当API信息定义冲突时，merge保留原有API信息 - override：当API信息定义冲突时，override会覆盖原有API信息
	ApiMode *def.MultiPart `json:"api_mode,omitempty"`

	// 导入Api的请求体，json或yaml格式的文件
	FileName *def.FilePart `json:"file_name"`
}

func (o ImportApiDefinitionsAsyncRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportApiDefinitionsAsyncRequestBody struct{}"
	}

	return strings.Join([]string{"ImportApiDefinitionsAsyncRequestBody", string(data)}, " ")
}

func (o *ImportApiDefinitionsAsyncRequestBody) UnmarshalJSON(b []byte) error {
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

type ImportApiDefinitionsAsyncRequestBodyExtendMode struct {
	value string
}

type ImportApiDefinitionsAsyncRequestBodyExtendModeEnum struct {
	MERGE    ImportApiDefinitionsAsyncRequestBodyExtendMode
	OVERRIDE ImportApiDefinitionsAsyncRequestBodyExtendMode
}

func GetImportApiDefinitionsAsyncRequestBodyExtendModeEnum() ImportApiDefinitionsAsyncRequestBodyExtendModeEnum {
	return ImportApiDefinitionsAsyncRequestBodyExtendModeEnum{
		MERGE: ImportApiDefinitionsAsyncRequestBodyExtendMode{
			value: "merge",
		},
		OVERRIDE: ImportApiDefinitionsAsyncRequestBodyExtendMode{
			value: "override",
		},
	}
}

func (c ImportApiDefinitionsAsyncRequestBodyExtendMode) Value() string {
	return c.value
}

func (c ImportApiDefinitionsAsyncRequestBodyExtendMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportApiDefinitionsAsyncRequestBodyExtendMode) UnmarshalJSON(b []byte) error {
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

type ImportApiDefinitionsAsyncRequestBodyApiMode struct {
	value string
}

type ImportApiDefinitionsAsyncRequestBodyApiModeEnum struct {
	MERGE    ImportApiDefinitionsAsyncRequestBodyApiMode
	OVERRIDE ImportApiDefinitionsAsyncRequestBodyApiMode
}

func GetImportApiDefinitionsAsyncRequestBodyApiModeEnum() ImportApiDefinitionsAsyncRequestBodyApiModeEnum {
	return ImportApiDefinitionsAsyncRequestBodyApiModeEnum{
		MERGE: ImportApiDefinitionsAsyncRequestBodyApiMode{
			value: "merge",
		},
		OVERRIDE: ImportApiDefinitionsAsyncRequestBodyApiMode{
			value: "override",
		},
	}
}

func (c ImportApiDefinitionsAsyncRequestBodyApiMode) Value() string {
	return c.value
}

func (c ImportApiDefinitionsAsyncRequestBodyApiMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportApiDefinitionsAsyncRequestBodyApiMode) UnmarshalJSON(b []byte) error {
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
