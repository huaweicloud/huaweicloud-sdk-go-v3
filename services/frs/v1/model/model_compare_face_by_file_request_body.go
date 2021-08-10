package model

import (
	"encoding/json"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"strings"
)

type CompareFaceByFileRequestBody struct {
	// 本地图片文件，图片不能超过8MB，建议小于1MB。上传文件时，请求格式为multipart。
	Image1File *def.FilePart `json:"image1_file"`

	// 本地图片文件，图片不能超过8MB，建议小于1MB。上传文件时，请求格式为multipart。
	Image2File *def.FilePart `json:"image2_file"`
}

func (o CompareFaceByFileRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CompareFaceByFileRequestBody struct{}"
	}

	return strings.Join([]string{"CompareFaceByFileRequestBody", string(data)}, " ")
}
