package model

import (
	"encoding/json"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"strings"
)

type DetectLiveFaceByFileRequestBody struct {
	// 本地图片文件。上传文件时，请求格式为multipart。
	ImageFile *def.FilePart `json:"image_file"`
}

func (o DetectLiveFaceByFileRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectLiveFaceByFileRequestBody struct{}"
	}

	return strings.Join([]string{"DetectLiveFaceByFileRequestBody", string(data)}, " ")
}
