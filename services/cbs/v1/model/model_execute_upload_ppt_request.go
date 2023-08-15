package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteUploadPptRequest Request Object
type ExecuteUploadPptRequest struct {

	// 视频id
	VideoId string `json:"video_id"`

	Body *ExecuteUploadPptRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ExecuteUploadPptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteUploadPptRequest struct{}"
	}

	return strings.Join([]string{"ExecuteUploadPptRequest", string(data)}, " ")
}
