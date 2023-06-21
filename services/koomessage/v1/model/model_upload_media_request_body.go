package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

	"strings"
)

type UploadMediaRequestBody struct {

	// 图片资源。  > 文件格式与文件名后缀需保持一致，请勿修改原始文件后缀，否则导致资源上传失败。
	File *def.FilePart `json:"file"`

	// 文件类型。 - PUB_LOGO：上传服务号LOGO - BG_IMAGE：上传服务号主页背景图 - FASTAPP_LOGO：上传快应用LOGO - OTHER：上传授权证明和营业执照等
	FileType *def.MultiPart `json:"file_type"`
}

func (o UploadMediaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadMediaRequestBody struct{}"
	}

	return strings.Join([]string{"UploadMediaRequestBody", string(data)}, " ")
}

func (o *UploadMediaRequestBody) UnmarshalJSON(b []byte) error {
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
