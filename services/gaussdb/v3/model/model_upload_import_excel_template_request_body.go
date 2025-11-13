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

type UploadImportExcelTemplateRequestBody struct {

	// **参数解释**：  具体选择哪一种模板进行下载。  **约束限制**：  不涉及。  **取值范围**：  import_async: Excel导入文件类型。  **默认取值**：  不涉及。
	TemplateType *def.MultiPart `json:"template_type"`

	// **参数解释**：  Excel文件上传。  **约束限制**：  Excel文件。  **取值范围**：  .xlsx文件。  **默认取值**：  不涉及。
	File *def.FilePart `json:"file"`

	// **参数解释**：  判断是否是实例级同步。  **约束限制**：  不涉及。  **取值范围**：  - true：实例级同步。 - false: 非实例级同步。  **默认取值**：  false。
	IsInstanceLevel *def.MultiPart `json:"is_instance_level,omitempty"`

	// **参数解释**：  用户选中的数据库名，用英文\",\"隔开。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SelectedDbs *def.MultiPart `json:"selected_dbs"`

	// **参数解释**：  是否支持标配符。  **约束限制**：  不涉及。  **取值范围**：  - true: 支持标配符。 - false: 不支持标配符。  **默认取值**：  不涉及。
	IsSupportRegexp *def.MultiPart `json:"is_support_regexp,omitempty"`
}

func (o UploadImportExcelTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadImportExcelTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"UploadImportExcelTemplateRequestBody", string(data)}, " ")
}

func (o *UploadImportExcelTemplateRequestBody) UnmarshalJSON(b []byte) error {
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

type UploadImportExcelTemplateRequestBodyIsInstanceLevel struct {
	value string
}

type UploadImportExcelTemplateRequestBodyIsInstanceLevelEnum struct {
	TRUE  UploadImportExcelTemplateRequestBodyIsInstanceLevel
	FALSE UploadImportExcelTemplateRequestBodyIsInstanceLevel
}

func GetUploadImportExcelTemplateRequestBodyIsInstanceLevelEnum() UploadImportExcelTemplateRequestBodyIsInstanceLevelEnum {
	return UploadImportExcelTemplateRequestBodyIsInstanceLevelEnum{
		TRUE: UploadImportExcelTemplateRequestBodyIsInstanceLevel{
			value: "true",
		},
		FALSE: UploadImportExcelTemplateRequestBodyIsInstanceLevel{
			value: "false",
		},
	}
}

func (c UploadImportExcelTemplateRequestBodyIsInstanceLevel) Value() string {
	return c.value
}

func (c UploadImportExcelTemplateRequestBodyIsInstanceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UploadImportExcelTemplateRequestBodyIsInstanceLevel) UnmarshalJSON(b []byte) error {
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

type UploadImportExcelTemplateRequestBodyIsSupportRegexp struct {
	value string
}

type UploadImportExcelTemplateRequestBodyIsSupportRegexpEnum struct {
	TRUE  UploadImportExcelTemplateRequestBodyIsSupportRegexp
	FALSE UploadImportExcelTemplateRequestBodyIsSupportRegexp
}

func GetUploadImportExcelTemplateRequestBodyIsSupportRegexpEnum() UploadImportExcelTemplateRequestBodyIsSupportRegexpEnum {
	return UploadImportExcelTemplateRequestBodyIsSupportRegexpEnum{
		TRUE: UploadImportExcelTemplateRequestBodyIsSupportRegexp{
			value: "true",
		},
		FALSE: UploadImportExcelTemplateRequestBodyIsSupportRegexp{
			value: "false",
		},
	}
}

func (c UploadImportExcelTemplateRequestBodyIsSupportRegexp) Value() string {
	return c.value
}

func (c UploadImportExcelTemplateRequestBodyIsSupportRegexp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UploadImportExcelTemplateRequestBodyIsSupportRegexp) UnmarshalJSON(b []byte) error {
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
