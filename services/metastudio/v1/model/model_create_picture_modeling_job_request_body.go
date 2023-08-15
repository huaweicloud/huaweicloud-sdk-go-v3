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

type CreatePictureModelingJobRequestBody struct {

	// 照片文件。 > 只能上传jpg/jpeg/png格式文件， 最大分辨率为3840*2160
	File *def.FilePart `json:"file"`

	// 数字人风格ID。 * system_male_001：男性风格01 * system_female_001：女性风格01 * system_male_002：男性风格02 * system_female_002：女性风格02
	StyleId *def.MultiPart `json:"style_id"`

	// 数字人模型名称，首次创建时使用。
	Name *def.MultiPart `json:"name"`

	// 照片建模任务结束的回调地址。
	NotifyUrl *def.MultiPart `json:"notify_url,omitempty"`
}

func (o CreatePictureModelingJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePictureModelingJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePictureModelingJobRequestBody", string(data)}, " ")
}

func (o *CreatePictureModelingJobRequestBody) UnmarshalJSON(b []byte) error {
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
