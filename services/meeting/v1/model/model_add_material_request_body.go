package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

	"strings"
)

type AddMaterialRequestBody struct {
	// 素材文件 - 只能上传jpg/jpeg/png格式文件，分辨率比率16:9，最大分辨率为3840*2160（推荐） - 请先命名完图片名称再上传
	File *def.FilePart `json:"file"`
}

func (o AddMaterialRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMaterialRequestBody struct{}"
	}

	return strings.Join([]string{"AddMaterialRequestBody", string(data)}, " ")
}

func (o *AddMaterialRequestBody) UnmarshalJSON(b []byte) error {
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
