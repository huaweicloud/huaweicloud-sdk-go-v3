package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"encoding/json"
	"os"
	"reflect"

	"strings"
)

type ImportLiveDataApiDefinitionsV2RequestBody struct {
	// 扩展信息导入模式 - merge：当扩展信息定义冲突时，merge保留原有扩展信息 - override：当扩展信息定义冲突时，override会覆盖原有扩展信息
	ExtendMode *def.MultiPart `json:"extend_mode,omitempty"`

	// 导入模式 - merge：当API信息定义冲突时，merge保留原有API信息 - override：当API信息定义冲突时，override会覆盖原有API信息
	ApiMode *def.MultiPart `json:"api_mode,omitempty"`

	// 导入自定义后端API的请求体，json或yaml格式的文件
	FileName *def.FilePart `json:"file_name"`
}

func (o ImportLiveDataApiDefinitionsV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportLiveDataApiDefinitionsV2RequestBody struct{}"
	}

	return strings.Join([]string{"ImportLiveDataApiDefinitionsV2RequestBody", string(data)}, " ")
}

func (o *ImportLiveDataApiDefinitionsV2RequestBody) UnmarshalJSON(b []byte) error {
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
		}
	}
	return nil
}

type ImportLiveDataApiDefinitionsV2RequestBodyExtendMode struct {
	value string
}

type ImportLiveDataApiDefinitionsV2RequestBodyExtendModeEnum struct {
	MERGE    ImportLiveDataApiDefinitionsV2RequestBodyExtendMode
	OVERRIDE ImportLiveDataApiDefinitionsV2RequestBodyExtendMode
}

func GetImportLiveDataApiDefinitionsV2RequestBodyExtendModeEnum() ImportLiveDataApiDefinitionsV2RequestBodyExtendModeEnum {
	return ImportLiveDataApiDefinitionsV2RequestBodyExtendModeEnum{
		MERGE: ImportLiveDataApiDefinitionsV2RequestBodyExtendMode{
			value: "merge",
		},
		OVERRIDE: ImportLiveDataApiDefinitionsV2RequestBodyExtendMode{
			value: "override",
		},
	}
}

func (c ImportLiveDataApiDefinitionsV2RequestBodyExtendMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportLiveDataApiDefinitionsV2RequestBodyExtendMode) UnmarshalJSON(b []byte) error {
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

type ImportLiveDataApiDefinitionsV2RequestBodyApiMode struct {
	value string
}

type ImportLiveDataApiDefinitionsV2RequestBodyApiModeEnum struct {
	MERGE    ImportLiveDataApiDefinitionsV2RequestBodyApiMode
	OVERRIDE ImportLiveDataApiDefinitionsV2RequestBodyApiMode
}

func GetImportLiveDataApiDefinitionsV2RequestBodyApiModeEnum() ImportLiveDataApiDefinitionsV2RequestBodyApiModeEnum {
	return ImportLiveDataApiDefinitionsV2RequestBodyApiModeEnum{
		MERGE: ImportLiveDataApiDefinitionsV2RequestBodyApiMode{
			value: "merge",
		},
		OVERRIDE: ImportLiveDataApiDefinitionsV2RequestBodyApiMode{
			value: "override",
		},
	}
}

func (c ImportLiveDataApiDefinitionsV2RequestBodyApiMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportLiveDataApiDefinitionsV2RequestBodyApiMode) UnmarshalJSON(b []byte) error {
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
