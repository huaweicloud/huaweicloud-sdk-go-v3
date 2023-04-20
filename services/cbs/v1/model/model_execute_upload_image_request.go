package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteUploadImageRequest struct {

	// 视频id
	VideoId string `json:"video_id"`

	// 图片名
	Name string `json:"name"`

	Body *ExecuteUploadImageRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ExecuteUploadImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteUploadImageRequest struct{}"
	}

	return strings.Join([]string{"ExecuteUploadImageRequest", string(data)}, " ")
}
