package model

import (
	"encoding/json"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"os"
	"reflect"

	"strings"
)

type CreateImageWatermarkRequestBody struct {
	// 要添加水印的图片文件，添加的图片短边尺寸需要超过512像素。
	File *def.FilePart `json:"file"`

	// 嵌入暗水印的内容，长度不超过32个字符。当前仅支持数字及英文大小写。
	BlindWatermark *def.MultiPart `json:"blind_watermark"`
}

func (o CreateImageWatermarkRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateImageWatermarkRequestBody struct{}"
	}

	return strings.Join([]string{"CreateImageWatermarkRequestBody", string(data)}, " ")
}

func (o *CreateImageWatermarkRequestBody) UnmarshalJSON(b []byte) error {
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
