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

// PostImagesReq 上传图片请求
type PostImagesReq struct {

	// 图片类型： 0：背景 最大 1920*1080 2：图标  最大1920*1080 图片格式：jpg，png
	Type *def.MultiPart `json:"type"`

	Name *def.MultiPart `json:"name"`

	// 图片文件
	File *def.FilePart `json:"file"`

	// 横竖屏模式 0：横屏（默认） 1：竖屏
	ResolutionType *def.MultiPart `json:"resolution_type,omitempty"`
}

func (o PostImagesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostImagesReq struct{}"
	}

	return strings.Join([]string{"PostImagesReq", string(data)}, " ")
}

func (o *PostImagesReq) UnmarshalJSON(b []byte) error {
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
