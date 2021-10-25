package model

import (
	"encoding/json"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"os"
	"reflect"

	"strings"
)

type UploadProtocolMappingsRequestBody struct {
	// 上传协议映射文件。当前仅支持xlsx/xls文件格式，且文件最大行数为10000行。
	File *def.FilePart `json:"file"`
}

func (o UploadProtocolMappingsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UploadProtocolMappingsRequestBody struct{}"
	}

	return strings.Join([]string{"UploadProtocolMappingsRequestBody", string(data)}, " ")
}

func (o *UploadProtocolMappingsRequestBody) UnmarshalJSON(b []byte) error {
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
