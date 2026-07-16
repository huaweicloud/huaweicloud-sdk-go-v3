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

type BatchBindApiKeyRequest struct {

	// **参数解释：** 请求批量绑定的api-key的id数组。 **约束限制：** 请求批量绑定api-key的id个数不超过10个。
	KeyIds *def.MultiPart `json:"key_ids"`
}

func (o BatchBindApiKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBindApiKeyRequest struct{}"
	}

	return strings.Join([]string{"BatchBindApiKeyRequest", string(data)}, " ")
}

func (o *BatchBindApiKeyRequest) UnmarshalJSON(b []byte) error {
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
