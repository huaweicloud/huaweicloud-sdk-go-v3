package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/def"
	"encoding/json"
	"os"
	"reflect"

	"strings"
)

type CompareFaceByFileRequestBody struct {
	// 本地图片文件，图片不能超过8MB。上传文件时，请求格式为multipart。
	Image1File *def.FilePart `json:"image1_file"`

	// 本地图片文件，图片不能超过8MB。上传文件时，请求格式为multipart。
	Image2File *def.FilePart `json:"image2_file"`
}

func (o CompareFaceByFileRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareFaceByFileRequestBody struct{}"
	}

	return strings.Join([]string{"CompareFaceByFileRequestBody", string(data)}, " ")
}

func (o *CompareFaceByFileRequestBody) UnmarshalJSON(b []byte) error {
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
