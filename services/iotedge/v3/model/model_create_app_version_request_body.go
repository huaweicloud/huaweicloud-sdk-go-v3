package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

	"strings"
)

type CreateAppVersionRequestBody struct {

	// chart包。当前仅支持tgz文件格式。
	Chart *def.FilePart `json:"chart"`

	// 应用包含的镜像列表
	Images *def.MultiPart `json:"images,omitempty"`
}

func (o CreateAppVersionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppVersionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAppVersionRequestBody", string(data)}, " ")
}

func (o *CreateAppVersionRequestBody) UnmarshalJSON(b []byte) error {
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
