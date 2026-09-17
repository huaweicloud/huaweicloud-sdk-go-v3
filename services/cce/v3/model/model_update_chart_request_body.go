package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

	"strings"
)

type UpdateChartRequestBody struct {

	// **参数解释：** 上传模板的配置参数，示例如下：\"{\\\"override\\\":true,\\\"skip_lint\\\":true,\\\"source\\\":\\\"package\\\"}\"。 **约束限制：** 不涉及 **取值范围：** - skip_lint：是否验证上传的模板 - override：是否覆盖已存在的模板 - visible：模板是否可见  **默认取值：** 不涉及
	Parameters *def.MultiPart `json:"parameters,omitempty"`

	// **参数解释：** 模板包文件。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Content *def.FilePart `json:"content"`
}

func (o UpdateChartRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateChartRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateChartRequestBody", string(data)}, " ")
}

func (o *UpdateChartRequestBody) UnmarshalJSON(b []byte) error {
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
