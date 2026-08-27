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

type ImportAlgorithmFileRequestBody struct {

	// **参数解释**： 算法最后更新时间，需要和算法的content_update_at字段一致。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	LastUpdateTime *def.MultiPart `json:"last_update_time"`

	// **参数解释**： 待上传文件。 **约束限制**： 需要是合法的zip压缩包。 **取值范围**： 20MB以内 **默认取值**： 不涉及
	File *def.FilePart `json:"file"`
}

func (o ImportAlgorithmFileRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportAlgorithmFileRequestBody struct{}"
	}

	return strings.Join([]string{"ImportAlgorithmFileRequestBody", string(data)}, " ")
}

func (o *ImportAlgorithmFileRequestBody) UnmarshalJSON(b []byte) error {
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
