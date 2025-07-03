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

type CreateAttachmentRequestBody struct {

	// **参数解释：** 需要上传的文件，以表单的形式上传。 **约束限制：** 内容为二进制binary文本。支持文件类型：.jpg,.png,.docx,.txt,.pdf，且文本大小不超10M。 **取值范围：** 不涉及 **默认取值：** 不涉及
	File *def.FilePart `json:"file"`
}

func (o CreateAttachmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAttachmentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAttachmentRequestBody", string(data)}, " ")
}

func (o *CreateAttachmentRequestBody) UnmarshalJSON(b []byte) error {
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
