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

type SaveAlgorithmFileRequestBody struct {

	// 文件存储路径
	FilePath *def.MultiPart `json:"file_path"`

	// 算法的最后更新时间
	LastUpdateTime *def.MultiPart `json:"last_update_time"`

	// **参数解释**： 待上传文件。 **约束限制**： 不涉及 **取值范围**： 5MB以内 **默认取值**： 不涉及
	File *def.FilePart `json:"file"`
}

func (o SaveAlgorithmFileRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveAlgorithmFileRequestBody struct{}"
	}

	return strings.Join([]string{"SaveAlgorithmFileRequestBody", string(data)}, " ")
}

func (o *SaveAlgorithmFileRequestBody) UnmarshalJSON(b []byte) error {
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
