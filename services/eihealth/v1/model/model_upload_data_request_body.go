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

type UploadDataRequestBody struct {

	// 文件流对象
	File *def.FilePart `json:"file"`

	// 目标文件夹
	TargetFolder *def.MultiPart `json:"target_folder,omitempty"`

	// 分段序号，表示第几个文件片段
	PartNumber *def.MultiPart `json:"part_number,omitempty"`

	// 分段总数，上传的文件总共分成了几个片段
	TotalPart *def.MultiPart `json:"total_part,omitempty"`

	// 分段上传任务id，除了第一个片段外，后续的片段都需要标识出任务id
	MultipartId *def.MultiPart `json:"multipart_id,omitempty"`

	// 文件名称
	FileName *def.MultiPart `json:"file_name,omitempty"`

	// MD5
	Md5 *def.MultiPart `json:"md5,omitempty"`
}

func (o UploadDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadDataRequestBody struct{}"
	}

	return strings.Join([]string{"UploadDataRequestBody", string(data)}, " ")
}

func (o *UploadDataRequestBody) UnmarshalJSON(b []byte) error {
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
