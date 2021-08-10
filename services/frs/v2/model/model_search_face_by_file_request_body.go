package model

import (
	"encoding/json"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"strings"
)

type SearchFaceByFileRequestBody struct {
	// 本地图片文件，图片不能超过8MB,建议小于1MB。上传文件时，请求格式为multipart。  必选，与image_url、image_base64、face_id四选一。
	ImageFile *def.FilePart `json:"image_file,omitempty"`

	// 返回查询到的最相似的N张人脸，N默认为10。
	TopN *def.MultiPart `json:"top_n,omitempty"`

	// 人脸相似度阈值，低于这个阈值则不返回，取值范围0~1，一般情况下建议取值0.93，默认为0。
	Threshold *def.MultiPart `json:"threshold,omitempty"`

	// 支持字段排序，参考[sort语法](zh-cn_topic_0130807047.xml)。
	Sort *def.MultiPart `json:"sort,omitempty"`

	// 过滤条件，参考[filter语法](zh-cn_topic_0130807048.xml)。
	Filter *def.MultiPart `json:"filter,omitempty"`

	// 指定返回的自定义字段。
	ReturnFields *def.MultiPart `json:"return_fields,omitempty"`
}

func (o SearchFaceByFileRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SearchFaceByFileRequestBody struct{}"
	}

	return strings.Join([]string{"SearchFaceByFileRequestBody", string(data)}, " ")
}
