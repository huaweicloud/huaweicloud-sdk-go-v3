package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"encoding/json"
	"os"
	"reflect"

	"strings"
)

type DetectLiveFaceByFileRequestBody struct {
	// 本地图片文件。上传文件时，请求格式为multipart。
	ImageFile *def.FilePart `json:"image_file"`
}

func (o DetectLiveFaceByFileRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveFaceByFileRequestBody struct{}"
	}

	return strings.Join([]string{"DetectLiveFaceByFileRequestBody", string(data)}, " ")
}

func (o *DetectLiveFaceByFileRequestBody) UnmarshalJSON(b []byte) error {
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
