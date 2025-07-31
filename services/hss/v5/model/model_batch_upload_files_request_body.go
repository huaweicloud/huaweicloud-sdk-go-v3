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

type BatchUploadFilesRequestBody struct {

	// **参数解释**: 上传文件类型 **约束限制**: 不涉及 **取值范围**: - dockerfile：Dockerfile文件。如果所有文件上传成功，接口返回文件名称和文件ID列表，服务保存上传的文件。如果批量上传的文件中存在上传失败的文件，接口返回上传成功和失败的文件信息，所有文件都不会被服务保存。 - k8s_yaml：k8s yaml文件。如果所有文件上传成功，接口返回文件名称和文件ID列表，服务保存上传的文件。如果批量上传的文件中存在上传失败的文件，接口返回上传成功和失败的文件信息，所有文件都不会被服务保存。  **默认取值**: 不涉及
	UploadType *def.MultiPart `json:"upload_type"`

	// **参数解释**: 上传的文件，支持批量上传 **约束限制**: 上传文件需要与upload_type的类型对应 **取值范围**: 单文件支持最大1M，10个。  **默认取值**: 不涉及
	Files *def.MultiPart `json:"files"`
}

func (o BatchUploadFilesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUploadFilesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUploadFilesRequestBody", string(data)}, " ")
}

func (o *BatchUploadFilesRequestBody) UnmarshalJSON(b []byte) error {
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
