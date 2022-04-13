package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/def"
	"encoding/json"
	"os"
	"reflect"

	"strings"
)

type ImportMqsInstanceTopicRequestBody struct {
	// 待导入的topic列表文件。
	UploadFileName *def.FilePart `json:"upload_file_name"`
}

func (o ImportMqsInstanceTopicRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportMqsInstanceTopicRequestBody struct{}"
	}

	return strings.Join([]string{"ImportMqsInstanceTopicRequestBody", string(data)}, " ")
}

func (o *ImportMqsInstanceTopicRequestBody) UnmarshalJSON(b []byte) error {
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
