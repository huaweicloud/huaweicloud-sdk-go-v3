package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"encoding/json"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

	"strings"
)

type ShowImageWatermarkRequestBody struct {
	// 待提取暗水印的图片文件。
	File *def.FilePart `json:"file"`

	// 指定待提取水印的长度，mark_len长度大于0，小于32。设置后可以提升水印提取性能
	MarkLen *def.MultiPart `json:"mark_len,omitempty"`
}

func (o ShowImageWatermarkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkRequestBody struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkRequestBody", string(data)}, " ")
}

func (o *ShowImageWatermarkRequestBody) UnmarshalJSON(b []byte) error {
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
