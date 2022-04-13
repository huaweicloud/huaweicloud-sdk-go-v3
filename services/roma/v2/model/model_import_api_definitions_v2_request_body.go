package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/def"
	"encoding/json"
	"os"
	"reflect"

	"strings"
)

type ImportApiDefinitionsV2RequestBody struct {
	// 是否创建新分组
	IsCreateGroup *def.MultiPart `json:"is_create_group,omitempty"`

	// API分组编号。  当is_create_group=false时为必填
	GroupId *def.MultiPart `json:"group_id,omitempty"`

	// 应用编号。  当is_create_group=false且使用集成应用分组时必填
	AppId *def.MultiPart `json:"app_id,omitempty"`

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

func (o ImportApiDefinitionsV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportApiDefinitionsV2RequestBody struct{}"
	}

	return strings.Join([]string{"ImportApiDefinitionsV2RequestBody", string(data)}, " ")
}

func (o *ImportApiDefinitionsV2RequestBody) UnmarshalJSON(b []byte) error {
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

type ImportApiDefinitionsV2RequestBodyExtendMode struct {
	value string
}

type ImportApiDefinitionsV2RequestBodyExtendModeEnum struct {
	MERGE    ImportApiDefinitionsV2RequestBodyExtendMode
	OVERRIDE ImportApiDefinitionsV2RequestBodyExtendMode
}

func GetImportApiDefinitionsV2RequestBodyExtendModeEnum() ImportApiDefinitionsV2RequestBodyExtendModeEnum {
	return ImportApiDefinitionsV2RequestBodyExtendModeEnum{
		MERGE: ImportApiDefinitionsV2RequestBodyExtendMode{
			value: "merge",
		},
		OVERRIDE: ImportApiDefinitionsV2RequestBodyExtendMode{
			value: "override",
		},
	}
}

func (c ImportApiDefinitionsV2RequestBodyExtendMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportApiDefinitionsV2RequestBodyExtendMode) UnmarshalJSON(b []byte) error {
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

type ImportApiDefinitionsV2RequestBodyApiMode struct {
	value string
}

type ImportApiDefinitionsV2RequestBodyApiModeEnum struct {
	MERGE    ImportApiDefinitionsV2RequestBodyApiMode
	OVERRIDE ImportApiDefinitionsV2RequestBodyApiMode
}

func GetImportApiDefinitionsV2RequestBodyApiModeEnum() ImportApiDefinitionsV2RequestBodyApiModeEnum {
	return ImportApiDefinitionsV2RequestBodyApiModeEnum{
		MERGE: ImportApiDefinitionsV2RequestBodyApiMode{
			value: "merge",
		},
		OVERRIDE: ImportApiDefinitionsV2RequestBodyApiMode{
			value: "override",
		},
	}
}

func (c ImportApiDefinitionsV2RequestBodyApiMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportApiDefinitionsV2RequestBodyApiMode) UnmarshalJSON(b []byte) error {
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
