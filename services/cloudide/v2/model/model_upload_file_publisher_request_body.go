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

type UploadFilePublisherRequestBody struct {

	// 文件
	File *def.FilePart `json:"file"`

	// 传用户token时，此字段为必传项
	PublisherId *def.MultiPart `json:"publisher_id,omitempty"`

	// 分片索引，第几个分片 取值范围：1-100
	ChunkIndex *def.MultiPart `json:"chunk_index"`

	// 是否合并已上传的分片（包含本次分片内容）,true
	Merge *def.MultiPart `json:"merge"`

	// 总分片数 0-100
	TotalChunkNum *def.MultiPart `json:"total_chunk_num"`

	// 父文件大小
	ParentFileSize *def.MultiPart `json:"parent_file_size"`

	// 父文件名称
	ParentFileName *def.MultiPart `json:"parent_file_name"`

	// 是否覆盖原有文件
	Override *def.MultiPart `json:"override"`

	// 文件分片的md5,用于校验文件分片是否完整
	ChunkMd5 *def.MultiPart `json:"chunk_md5"`

	// 父文件hash，用于校验合并后的文件
	ParentFileSha256 *def.MultiPart `json:"parent_file_sha256,omitempty"`

	// 上传任务的唯一标识，第一次上传分片时可不传
	TaskId *def.MultiPart `json:"task_id,omitempty"`
}

func (o UploadFilePublisherRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFilePublisherRequestBody struct{}"
	}

	return strings.Join([]string{"UploadFilePublisherRequestBody", string(data)}, " ")
}

func (o *UploadFilePublisherRequestBody) UnmarshalJSON(b []byte) error {
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
