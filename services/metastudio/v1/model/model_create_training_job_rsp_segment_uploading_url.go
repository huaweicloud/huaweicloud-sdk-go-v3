package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrainingJobRspSegmentUploadingUrl 分句上传任务的上传地址,create_type为segment时设置。
type CreateTrainingJobRspSegmentUploadingUrl struct {

	// 音频上传的地址。  通过该obs地址上传时，需设置content-type为audio/wav
	AudioUploadingUrl *[]string `json:"audio_uploading_url,omitempty"`

	// 文本上传的地址。  通过该obs地址上传时需设置content-type为text/plain
	TxtUploadingUrl *[]string `json:"txt_uploading_url,omitempty"`
}

func (o CreateTrainingJobRspSegmentUploadingUrl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainingJobRspSegmentUploadingUrl struct{}"
	}

	return strings.Join([]string{"CreateTrainingJobRspSegmentUploadingUrl", string(data)}, " ")
}
